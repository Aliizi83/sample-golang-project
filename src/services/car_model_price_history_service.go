package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelPriceHistoryService(cfg *config.Config, repository repositories.CarModelPriceHistoryRepository) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		BaseService: NewBaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory](cfg, repository),
	}
}

type CarModelPriceHistoryService struct {
	*BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory]
}