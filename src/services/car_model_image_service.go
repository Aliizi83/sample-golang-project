package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelImageService(cfg *config.Config, repository repositories.CarModelImageRepository) *CarModelImageService {
	return &CarModelImageService{
		BaseService: NewBaseService[models.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage](cfg, repository),
	}
}

type CarModelImageService struct {
	*BaseService[models.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage]
}