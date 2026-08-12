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

type CarModelPropertyHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelPropertyService *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	carModelPropertyRepository := dependencies.GetCarModelPropertyRepository(cfg)

	return &CarModelPropertyHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelPropertyService: services.NewCarModelPropertyService(cfg, carModelPropertyRepository),
	}
}

// CreateCarModelProperty godoc
// @Summary Create a carModelProperty
// @Description Create a carModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPropertyRequest true "Create a carModelProperty"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/ [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelProperty, dto.ToCarModelPropertyResponse, h.carModelPropertyService.Create)
}

// UpdateCarModelProperty godoc
// @Summary Update a carModelProperty
// @Description Update a carModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPropertyRequest true "Update a carModelProperty"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/{id} [put]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelProperty, dto.ToCarModelPropertyResponse, h.carModelPropertyService.Update)
}

// DeleteCarModelProperty godoc
// @Summary Delete a carModelProperty
// @Description Delete a carModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/{id} [delete]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelPropertyService.Delete)
}

// GetCarModelProperty godoc
// @Summary Get a carModelProperty
// @Description Get a carModelProperty
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelPropertyResponse} "carModelProperty response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/{id} [get]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelPropertyResponse, h.carModelPropertyService.GetById)
}

// GetCarModelProperties godoc
// @Summary Get CarModelProperties
// @Description Get CarModelProperties
// @Tags CarModelProperties
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelPropertyResponse]} "CarModelProperty response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-properties/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelPropertyResponse, h.carModelPropertyService.GetByFilter)
}
