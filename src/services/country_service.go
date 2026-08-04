package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateAndUpdateCountryRequest, dto.CreateAndUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	base := NewBaseService[models.Country, dto.CreateAndUpdateCountryRequest, dto.CreateAndUpdateCountryRequest, dto.CountryResponse](cfg)
	return &CountryService{base: base}
}

func (s *CountryService) Create(ctx context.Context, req dto.CreateAndUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CountryService) Update(ctx context.Context, id int, req dto.CreateAndUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}
