package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelPropertyService(cfg *config.Config, repository repositories.CarModelPropertyRepository) *CarModelPropertyService {
	return &CarModelPropertyService{
		BaseService: NewBaseService[models.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty](cfg, repository),
	}
}

type CarModelPropertyService struct {
	*BaseService[models.CarModelProperty, dto.CreateCarModelProperty, dto.UpdateCarModelProperty, dto.CarModelProperty]
}