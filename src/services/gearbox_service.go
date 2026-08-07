package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewGearboxService(cfg *config.Config, repository repositories.GearboxRepository) *GearboxService {
	return &GearboxService{
		BaseService: NewBaseService[models.Gearbox, dto.CreateGearbox, dto.UpdateGearbox, dto.Gearbox](cfg, repository),
	}
}

type GearboxService struct {
	*BaseService[models.Gearbox, dto.CreateGearbox, dto.UpdateGearbox, dto.Gearbox]
}