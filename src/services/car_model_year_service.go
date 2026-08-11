package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelYearService(cfg *config.Config, repository repositories.CarModelYearRepository) *CarModelYearService {
	return &CarModelYearService{
		BaseService: NewBaseService[models.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear](cfg, repository),
	}
}

type CarModelYearService struct {
	*BaseService[models.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear]
}