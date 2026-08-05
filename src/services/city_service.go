package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	dependencies "github.com/Aliizi83/sample-golang-project/src/dependenies"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		BaseService: NewBaseService[models.City, dto.CreateCity, dto.UpdateCity, dto.City](cfg, dependencies.GetCityRepository(cfg)),
	}
}

type CityService struct {
	*BaseService[models.City, dto.CreateCity, dto.UpdateCity, dto.City]
}
