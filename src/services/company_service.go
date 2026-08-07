package services

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCompanyService(cfg *config.Config, repository repositories.CompanyRepository) *CompanyService {
	return &CompanyService{
		BaseService: NewBaseService[models.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company](cfg, repository),
	}
}

type CompanyService struct {
	*BaseService[models.Company, dto.CreateCompany, dto.UpdateCompany, dto.Company]
}