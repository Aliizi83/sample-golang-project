package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"gorm.io/gorm"
)

type BaseService[TModel, TCreate, TUpdate, TResponse any] struct {
	Database *gorm.DB
	Logger   logging.Logger
}

func NewBaseService[TModel, TCreate, TUpdate, TResponse any](cfg *config.Config) *BaseService[TModel, TCreate, TUpdate, TResponse] {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	return &BaseService[TModel, TCreate, TUpdate, TResponse]{
		Database: db,
		Logger:   logger,
	}
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Create(ctx context.Context, req TCreate) (*TResponse, error) {
	response := new(TResponse)
	model, err := common.TypeConverter[TModel](req)
	if err != nil {
		return response, err
	}

	tx := s.Database.WithContext(ctx).Begin()
	err = tx.Create(&model).Error
	if err != nil {
		s.Logger.Error(err, logging.Postgres, logging.Insert, err.Error(), nil)
		tx.Rollback()
		return response, err
	}

	return common.TypeConverter[TResponse](model)
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Update(ctx context.Context, id int, req TUpdate) (*TResponse, error) {
	var model TModel
	response := new(TResponse)
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return response, nil
	}
	snakeCaseUpdateMap := make(map[string]any, len(*updateMap))

	for k, v := range *updateMap {
		snakeCaseUpdateMap[common.ToSnakeCase(k)] = v
	}

	snakeCaseUpdateMap["modified_at"] = sql.NullTime{Time: time.Now(), Valid: true}
	snakeCaseUpdateMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}

	tx := s.Database.WithContext(ctx).Begin()
	err = tx.Model(model).Where("id = ? AND deleted_by IS NULL").Updates(snakeCaseUpdateMap).Error

	if err != nil {
		s.Logger.Error(err, logging.Postgres, logging.Update, err.Error(), nil)
		tx.Rollback()
		return response, err
	}

	return response, nil
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Delete(ctx context.Context, id int) error {
	model := new(TModel)
	updateMap := map[string]any{
		"modified_at": sql.NullTime{Time: time.Now(), Valid: true},
		"modified_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
	}

	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Model(model).Where("id = ? AND deleted_by IS NULL").Updates(updateMap).Error
	if err != nil {
		s.Logger.Error(err, logging.Postgres, logging.Delete, err.Error(), nil)
		tx.Rollback()
		return err
	}

	return err
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) GetById(ctx context.Context, id int) (*TResponse, error) {
	model := new(TModel)

	err := s.Database.WithContext(ctx).Where("id = ? AND deleted_by IS NULL").First(model).Error
	if err != nil {
		s.Logger.Error(err, logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}

	return common.TypeConverter[TResponse](model)

}
