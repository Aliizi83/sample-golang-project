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

type CarModelColorHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelColorService *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	carModelColorRepository := dependencies.GetCarModelColorRepository(cfg)

	return &CarModelColorHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelColorService: services.NewCarModelColorService(cfg, carModelColorRepository),
	}
}

// CreateCarModelColor godoc
// @Summary Create a carModelColor
// @Description Create a carModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelColorRequest true "Create a carModelColor"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/ [post]
// @Security AuthBearer
func (h *CarModelColorHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelColor, dto.ToCarModelColorResponse, h.carModelColorService.Create)
}

// UpdateCarModelColor godoc
// @Summary Update a carModelColor
// @Description Update a carModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelColorRequest true "Update a carModelColor"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [put]
// @Security AuthBearer
func (h *CarModelColorHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelColor, dto.ToCarModelColorResponse, h.carModelColorService.Update)
}

// DeleteCarModelColor godoc
// @Summary Delete a carModelColor
// @Description Delete a carModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelColorService.Delete)
}

// GetCarModelColor godoc
// @Summary Get a carModelColor
// @Description Get a carModelColor
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelColorResponse} "carModelColor response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/{id} [get]
// @Security AuthBearer
func (h *CarModelColorHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelColorResponse, h.carModelColorService.GetById)
}

// GetCarModelColors godoc
// @Summary Get CarModelColors
// @Description Get CarModelColors
// @Tags CarModelColors
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelColorResponse]} "CarModelColor response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-colors/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelColorHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelColorResponse, h.carModelColorService.GetByFilter)
}
