package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarTypeService(cfg *config.Config, repository repositories.CarTypeRepository) *CarTypeService {
	return &CarTypeService{
		BaseService: NewBaseService[models.CarType, dto.CreateCarType, dto.UpdateCarType, dto.CarType](cfg, repository),
	}
}

type CarTypeService struct {
	*BaseService[models.CarType, dto.CreateCarType, dto.UpdateCarType, dto.CarType]
}