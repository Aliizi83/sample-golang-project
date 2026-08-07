package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewPropertyService(cfg *config.Config, repository repositories.PropertyRepository) *PropertyService {
	return &PropertyService{
		BaseService: NewBaseService[models.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property](cfg, repository),
	}
}

type PropertyService struct {
	*BaseService[models.Property, dto.CreateProperty, dto.UpdateProperty, dto.Property]
}