package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewPersianYearService(cfg *config.Config, repository repositories.PersianYearRepository) *PersianYearService {
	return &PersianYearService{
		BaseService: NewBaseService[models.PersianYear, dto.CreatePersianYear, dto.UpdatePersianYear, dto.PersianYear](cfg, repository),
	}
}

type PersianYearService struct {
	*BaseService[models.PersianYear, dto.CreatePersianYear, dto.UpdatePersianYear, dto.PersianYear]
}