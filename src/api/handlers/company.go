package handlers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/dependencies"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"

	_ "github.com/Aliizi83/sample-golang-project/src/api/helpers"
	_ "github.com/Aliizi83/sample-golang-project/src/domain/filters"
)

type CompanyHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	companyService *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	companyRepository := dependencies.GetCompanyRepository(cfg)

	return &CompanyHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		companyService: services.NewCompanyService(cfg, companyRepository),
	}
}

// CreateCompany godoc
// @Summary Create a company
// @Description Create a company
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body dto.CreateCompanyRequest true "Create a company"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/companies/ [post]
// @Security AuthBearer
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCompany, dto.ToCompanyResponse, h.companyService.Create)
}

// UpdateCompany godoc
// @Summary Update a company
// @Description Update a company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCompanyRequest true "Update a company"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CompanyResponse} "Company response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [put]
// @Security AuthBearer
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCompany, dto.ToCompanyResponse, h.companyService.Update)
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description Delete a company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [delete]
// @Security AuthBearer
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.companyService.Delete)
}

// GetCompany godoc
// @Summary Get a company
// @Description Get a company
// @Tags Companies
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CompanyResponse} "company response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/companies/{id} [get]
// @Security AuthBearer
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCompanyResponse, h.companyService.GetById)
}

// GetCompanies godoc
// @Summary Get Companies
// @Description Get Companies
// @Tags Companies
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CompanyResponse]} "Company response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/companies/get-by-filters [post]
// @Security AuthBearer
func (h *CompanyHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCompanyResponse, h.companyService.GetByFilter)
}
