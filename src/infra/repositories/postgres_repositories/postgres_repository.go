package postgres_repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? AND deleted_by IS NULL"

type BaseRepository[TModel any] struct {
	database *gorm.DB
	logger   logging.Logger
	preloads []db.PreloadEntity
}

func NewBaseRepository[TModel any](cfg *config.Config, preloads []db.PreloadEntity) *BaseRepository[TModel] {
	return &BaseRepository[TModel]{
		database: db.GetDB(),
		logger:   logging.NewLogger(cfg),
		preloads: preloads,
	}
}

func (r *BaseRepository[TModel]) Create(ctx context.Context, model TModel) (TModel, error) {
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Create(&model).Error; err != nil {
		modelName := common.GetTypeName[TModel]()
		r.logger.Error(err, logging.Postgres, logging.Insert, fmt.Sprintf("error while creating %s : %s ", modelName, err.Error()), nil)
		tx.Rollback()
		return model, err
	}

	tx.Commit()
	return model, nil
}

func (r *BaseRepository[TModel]) Update(ctx context.Context, id int, updateFields map[string]any) (TModel, error) {
	var model TModel
	tx := r.database.WithContext(ctx).Begin()
	snakeUpdateFields := make(map[string]any)
	for k, v := range updateFields {
		snakeUpdateFields[common.ToSnakeCase(k)] = v
	}

	snakeUpdateFields["modified_by"] = &sql.NullInt64{Int64: int64(r.getClaims(ctx, constants.UserIdKey).(float64)), Valid: true}
	snakeUpdateFields["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	if err := tx.Model(&model).Where(softDeleteExp, id).Updates(snakeUpdateFields).Error; err != nil {
		tx.Rollback()
		modelName := common.GetTypeName[TModel]()
		r.logger.Error(err, logging.Postgres, logging.Update, fmt.Sprintf("error while updating %s : %s ", modelName, err.Error()), nil)
		return model, err
	}
	tx.Commit()
	return model, nil
}
func (r *BaseRepository[TModel]) Delete(ctx context.Context, id int) error {
	var model TModel
	tx := r.database.WithContext(ctx).Begin()
	updateFields := make(map[string]any)

	updateFields["deleted_by"] = &sql.NullInt64{Int64: int64(r.getClaims(ctx, constants.UserIdKey).(float64)), Valid: true}
	updateFields["deleted_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	if err := tx.Model(&model).Delete(softDeleteExp, id).Updates(updateFields).Error; err != nil {
		tx.Rollback()
		modelName := common.GetTypeName[TModel]()
		r.logger.Error(err, logging.Postgres, logging.Update, fmt.Sprintf("error while deleting %s : %s ", modelName, err.Error()), nil)
		return err
	}
	tx.Commit()
	return nil
}
func (r *BaseRepository[TModel]) GetById(ctx context.Context, id int) (TModel, error) {
	var model TModel
	if err := r.database.Where(softDeleteExp, id).Where("id = ?", id).First(&model).Error; err != nil {
		modelName := common.GetTypeName[TModel]()
		r.logger.Error(err, logging.Postgres, logging.Insert, fmt.Sprintf("error while selecting %s : %s ", modelName, err.Error()), nil)
		return model, err
	}

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
