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

type CarModelHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelService *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	carModelRepository := dependencies.GetCarModelRepository(cfg)

	return &CarModelHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelService: services.NewCarModelService(cfg, carModelRepository),
	}
}

// CreateCarModel godoc
// @Summary Create a carModel
// @Description Create a carModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelRequest true "Create a carModel"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModel, dto.ToCarModelResponse, h.carModelService.Create)
}

// UpdateCarModel godoc
// @Summary Update a carModel
// @Description Update a carModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a carModel"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModel, dto.ToCarModelResponse, h.carModelService.Update)
}

// DeleteCarModel godoc
// @Summary Delete a carModel
// @Description Delete a carModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelService.Delete)
}

// GetCarModel godoc
// @Summary Get a carModel
// @Description Get a carModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelResponse} "carModel response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelResponse, h.carModelService.GetById)
}

// GetCarModels godoc
// @Summary Get CarModels
// @Description Get CarModels
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelResponse]} "CarModel response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-models/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelResponse, h.carModelService.GetByFilter)
}
