package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCityService(cfg *config.Config, repository repositories.CityRepository) *CityService {
	return &CityService{
		BaseService: NewBaseService[models.City, dto.CreateCity, dto.UpdateCity, dto.City](cfg, repository),
	}
}

type CityService struct {
	*BaseService[models.City, dto.CreateCity, dto.UpdateCity, dto.City]
}
