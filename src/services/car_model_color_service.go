package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelColorService(cfg *config.Config, repository repositories.CarModelColorRepository) *CarModelColorService {
	return &CarModelColorService{
		BaseService: NewBaseService[models.CarModelColor, dto.CreateCarModelColor, dto.UpdateCarModelColor, dto.CarModelColor](cfg, repository),
	}
}

type CarModelColorService struct {
	*BaseService[models.CarModelColor, dto.CreateCarModelColor, dto.UpdateCarModelColor, dto.CarModelColor]
}