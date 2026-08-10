package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewColorService(cfg *config.Config, repository repositories.ColorRepository) *ColorService {
	return &ColorService{
		BaseService: NewBaseService[models.Color, dto.CreateColor, dto.UpdateColor, dto.Color](cfg, repository),
	}
}

type ColorService struct {
	*BaseService[models.Color, dto.CreateColor, dto.UpdateColor, dto.Color]
}