package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
)

type BaseService[TModel models.Identifiable, TCreate, TUpdate, TResponse any] struct {
	logger     logging.Logger
	repository repositories.BaseRepository[TModel]
}

func NewBaseService[TModel models.Identifiable, TCreate, TUpdate, TResponse any](cfg *config.Config, repository repositories.BaseRepository[TModel]) *BaseService[TModel, TCreate, TUpdate, TResponse] {
	logger := logging.NewLogger(cfg)
	return &BaseService[TModel, TCreate, TUpdate, TResponse]{
		repository: repository,
		logger:     logger,
	}
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Create(ctx context.Context, req TCreate) (TResponse, error) {
	var response TResponse
	model, err := common.TypeConverter[TModel](req)
	if err != nil {
		return response, err
	}

	model, err = s.repository.Create(ctx, model)
	if err != nil {
		s.logger.Error(err, logging.Postgres, logging.Insert, err.Error(), nil)
		return response, err
	}

	response, _ = common.TypeConverter[TResponse](model)
	return response, nil
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Update(ctx context.Context, id int, req TUpdate) (TResponse, error) {
	var response TResponse
	updateMap, err := common.TypeConverter[map[string]any](req)
	if err != nil {
		s.logger.Error(err, logging.General, logging.TypeConverting, err.Error(), nil)
		return response, err
	}

	model, err := s.repository.Update(ctx, id, updateMap)

	if err != nil {
		return response, err
	}

	response, err = common.TypeConverter[TResponse](model)
	if err != nil {
		s.logger.Error(err, logging.General, logging.TypeConverting, err.Error(), nil)
		return response, err
	}

	return s.GetById(ctx, id)
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) GetById(ctx context.Context, id int) (TResponse, error) {
	var response TResponse
	model, err := s.repository.GetById(ctx, id)
	if err != nil {
		return response, err
	}

	return common.TypeConverter[TResponse](model)
}

func (s *BaseService[TModel, TCreate, TUpdate, TResponse]) GetByFilter(ctx context.Context, req filters.PaginationInputWithFilter) (*filters.PagedList[TResponse], error) {
	var response *filters.PagedList[TResponse]
	count, models, err := s.repository.GetByFilter(ctx, req)
	if err != nil {
		return response, err
	}

	return filters.Paginate[TModel, TResponse](count, models, req.PageNumber, int64(req.PageSize))
}
