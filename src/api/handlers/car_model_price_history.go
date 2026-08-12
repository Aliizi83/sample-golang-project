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

type CarModelPriceHistoryHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelPriceHistoryService *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	carModelPriceHistoryRepository := dependencies.GetCarModelPriceHistoryRepository(cfg)

	return &CarModelPriceHistoryHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelPriceHistoryService: services.NewCarModelPriceHistoryService(cfg, carModelPriceHistoryRepository),
	}
}

// CreateCarModelPriceHistory godoc
// @Summary Create a carModelPriceHistory
// @Description Create a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "Create a carModelPriceHistory"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelPriceHistory, dto.ToCarModelPriceHistoryResponse, h.carModelPriceHistoryService.Create)
}

// UpdateCarModelPriceHistory godoc
// @Summary Update a carModelPriceHistory
// @Description Update a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a carModelPriceHistory"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelPriceHistory, dto.ToCarModelPriceHistoryResponse, h.carModelPriceHistoryService.Update)
}

// DeleteCarModelPriceHistory godoc
// @Summary Delete a carModelPriceHistory
// @Description Delete a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelPriceHistoryService.Delete)
}

// GetCarModelPriceHistory godoc
// @Summary Get a carModelPriceHistory
// @Description Get a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "carModelPriceHistory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelPriceHistoryResponse, h.carModelPriceHistoryService.GetById)
}

// GetCarModelPriceHistories godoc
// @Summary Get CarModelPriceHistories
// @Description Get CarModelPriceHistories
// @Tags CarModelPriceHistories
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelPriceHistoryResponse]} "CarModelPriceHistory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-price-histories/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelPriceHistoryResponse, h.carModelPriceHistoryService.GetByFilter)
}
