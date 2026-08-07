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

type CityHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	cityService *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	repository := dependencies.GetCityRepository(cfg)
	return &CityHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		cityService: services.NewCityService(cfg, repository),
	}
}

// CreateCity godoc
// @Summary Create a city
// @Description Create a city
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body dto.CreateCityRequest true "Create a city"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCity, dto.ToCityResponse, h.cityService.Create)
}

// UpdateCity godoc
// @Summary Update a city
// @Description Update a city
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCityRequest true "Update a city"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCity, dto.ToCityResponse, h.cityService.Update)
}

// DeleteCity godoc
// @Summary Delete a city
// @Description Delete a city
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	Delete(c, h.cityService.Delete)
}

// GetCity godoc
// @Summary Get a city
// @Description Get a city
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CityResponse} "city response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCityResponse, h.cityService.GetById)
}

// GetCities godoc
// @Summary Get Cities
// @Description Get Cities
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CityResponse]} "City response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/cities/get-by-filters [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCityResponse, h.cityService.GetByFilter)
}
