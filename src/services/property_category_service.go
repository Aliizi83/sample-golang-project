package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewPropertyCategoryService(cfg *config.Config, repository repositories.PropertyCategoryRepository) *PropertyCategoryService {
	return &PropertyCategoryService{
		BaseService: NewBaseService[models.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory](cfg, repository),
	}
}

type PropertyCategoryService struct {
	*BaseService[models.PropertyCategory, dto.CreatePropertyCategory, dto.UpdatePropertyCategory, dto.PropertyCategory]
}