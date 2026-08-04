package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

type CountryService struct {
	base *BaseService[models.Country, dto.Name, dto.Name, dto.Country]
}

func NewCountryService(cfg *config.Config, repository repositories.CountryRepository) *CountryService {
	base := NewBaseService[models.Country, dto.Name, dto.Name, dto.Country](cfg, repository)
	return &CountryService{base: base}
}

func (s *CountryService) Create(ctx context.Context, req dto.Name) (dto.Country, error) {
	return s.base.Create(ctx, req)
}

func (s *CountryService) Update(ctx context.Context, id int, req dto.Name) (dto.Country, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetById(ctx context.Context, id int) (dto.Country, error) {
	return s.base.GetById(ctx, id)
}

func (s *CountryService) GetByFilter(ctx context.Context, req filters.PaginationInputWithFilter) (*filters.PagedList[dto.Country], error) {
	return s.base.GetByFilter(ctx, req)
}
