package handlers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	_ "github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	dependencies "github.com/Aliizi83/sample-golang-project/src/dependenies"
	_ "github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	cfg            *config.Config
	logger         logging.Logger
	countryService services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	logger := logging.NewLogger(cfg)
	repository := dependencies.GetCountryRepository(cfg)
	countryService := services.NewCountryService(cfg, repository)

	return &CountryHandler{
		cfg:            cfg,
		logger:         logger,
		countryService: *countryService,
	}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.CreateUpdateCountryRequest true "Create a country"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateUpdateCountry, dto.ToCountryResponse, h.countryService.Create)
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a country"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	Update(c, dto.ToCreateUpdateCountry, dto.ToCountryResponse, h.countryService.Update)
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	Delete(c, h.countryService.Delete)
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCountryResponse, h.countryService.GetById)
}

// GetCountries godoc
// @Summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CountryResponse]} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/get-by-filters [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCountryResponse, h.countryService.GetByFilter)
}
