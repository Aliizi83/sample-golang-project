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

type CarTypeHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carTypeService *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	carTypeRepository := dependencies.GetCarTypeRepository(cfg)

	return &CarTypeHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carTypeService: services.NewCarTypeService(cfg, carTypeRepository),
	}
}

// CreateCarType godoc
// @Summary Create a carType
// @Description Create a carType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a carType"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-types/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarType, dto.ToCarTypeResponse, h.carTypeService.Create)
}

// UpdateCarType godoc
// @Summary Update a carType
// @Description Update a carType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a carType"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarType, dto.ToCarTypeResponse, h.carTypeService.Update)
}

// DeleteCarType godoc
// @Summary Delete a carType
// @Description Delete a carType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.carTypeService.Delete)
}

// GetCarType godoc
// @Summary Get a carType
// @Description Get a carType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarTypeResponse} "carType response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarTypeResponse, h.carTypeService.GetById)
}

// GetCarTypes godoc
// @Summary Get CarTypes
// @Description Get CarTypes
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarTypeResponse]} "CarType response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-types/get-by-filters [post]
// @Security AuthBearer
func (h *CarTypeHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarTypeResponse, h.carTypeService.GetByFilter)
}
