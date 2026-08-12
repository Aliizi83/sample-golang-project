package postgres_repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/metrics"
	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? AND deleted_by IS NULL"

type BaseRepository[TModel models.Identifiable] struct {
	database  *gorm.DB
	logger    logging.Logger
	preloads  []db.PreloadEntity
	modelName string
}

func NewBaseRepository[TModel models.Identifiable](cfg *config.Config, preloads []db.PreloadEntity) *BaseRepository[TModel] {
	return &BaseRepository[TModel]{
		database:  db.GetDB(),
		logger:    logging.NewLogger(cfg),
		preloads:  preloads,
		modelName: common.GetTypeName[TModel](),
	}
}

func (r *BaseRepository[TModel]) Create(ctx context.Context, model TModel) (TModel, error) {
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Create(&model).Error; err != nil {
		modelName := common.GetTypeName[TModel]()
		r.logger.Error(err, logging.Postgres, logging.Insert, fmt.Sprintf("error while creating %s : %s ", modelName, err.Error()), nil)
		metrics.DbCall.WithLabelValues(modelName, "Create", "Failed")

		tx.Rollback()
		return model, err
	}
	tx.Commit()
	id := model.GetID()
	return r.GetById(ctx, int(id))
}

func (r *BaseRepository[TModel]) Update(ctx context.Context, id int, updateFields map[string]any) (TModel, error) {
	var model TModel
	tx := r.database.WithContext(ctx).Begin()
	snakeUpdateFields := make(map[string]any)
	for k, v := range updateFields {
		snakeUpdateFields[common.ToSnakeCase(k)] = v
	}

	userId := int64(r.getClaims(ctx, constants.UserIdKey).(float64))
	snakeUpdateFields["modified_by"] = userId
	snakeUpdateFields["modified_at"] = time.Now().UTC()

	if err := tx.Model(&model).Where(softDeleteExp, id).Updates(snakeUpdateFields).Error; err != nil {
		tx.Rollback()
		r.logger.Error(err, logging.Postgres, logging.Update, fmt.Sprintf("error while updating %s : %s ", r.modelName, err.Error()), nil)
		metrics.DbCall.WithLabelValues(r.modelName, "Update", "Failed")

		return model, err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(r.modelName, "Update", "Success")
	return r.GetById(ctx, int(id))
}
func (r *BaseRepository[TModel]) Delete(ctx context.Context, id int) error {
	var model TModel
	tx := r.database.WithContext(ctx).Begin()
	updateFields := make(map[string]any)

	userId := int64(r.getClaims(ctx, constants.UserIdKey).(float64))
	updateFields["deleted_by"] = userId
	updateFields["deleted_at"] = time.Now().UTC()

	if err := tx.Model(&model).Where(softDeleteExp, id).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		r.logger.Error(err, logging.Postgres, logging.Update, fmt.Sprintf("error while deleting %s : %s ", r.modelName, err.Error()), nil)
		metrics.DbCall.WithLabelValues(r.modelName, "Delete", "Failed")
		return err
	}
	tx.Commit()
	metrics.DbCall.WithLabelValues(r.modelName, "Delete", "Success")
	return nil
}
func (r *BaseRepository[TModel]) GetById(ctx context.Context, id int) (TModel, error) {
	var model TModel
	database := db.Preload(r.database, r.preloads)
	if err := database.Where(softDeleteExp, id).Where("id = ?", id).First(&model).Error; err != nil {
		r.logger.Error(err, logging.Postgres, logging.Insert, fmt.Sprintf("error while selecting %s : %s ", r.modelName, err.Error()), nil)
		metrics.DbCall.WithLabelValues(r.modelName, "GetById", "Failed")
		return model, err
	}

	metrics.DbCall.WithLabelValues(r.modelName, "GetById", "Success")
	return model, nil
}
func (r *BaseRepository[TModel]) GetByFilter(ctx context.Context, req filters.PaginationInputWithFilter) (int64, *[]TModel, error) {
	model := new(TModel)
	var items *[]TModel

	database := db.Preload(r.database, r.preloads)
	query := db.GenerateDynamicQuery[TModel](&req.DynamicFilter)
	sort := db.GenerateDynamicSort[TModel](&req.DynamicFilter)

	var totalRows int64 = 0

	database.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := database.
		Where(query).
		Offset(req.GetOffset()).
		Limit(req.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return 0, &[]TModel{}, err
	}

	return totalRows, items, err
}

func (r *BaseRepository[TModel]) getClaims(ctx context.Context, key string) any {
	claims := ctx.Value(constants.ClaimsKey).(map[string]any)
	value, ok := claims[key]
	if !ok {
		return nil
	}
	return value
}
