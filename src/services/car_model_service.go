package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelService(cfg *config.Config, repository repositories.CarModelRepository) *CarModelService {
	return &CarModelService{
		BaseService: NewBaseService[models.CarModel, dto.CreateCarModel, dto.UpdateCarModel, dto.CarModel](cfg, repository),
	}
}

type CarModelService struct {
	*BaseService[models.CarModel, dto.CreateCarModel, dto.UpdateCarModel, dto.CarModel]
}