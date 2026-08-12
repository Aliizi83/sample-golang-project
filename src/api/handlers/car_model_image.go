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

type CarModelImageHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelImageService *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	carModelImageRepository := dependencies.GetCarModelImageRepository(cfg)

	return &CarModelImageHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelImageService: services.NewCarModelImageService(cfg, carModelImageRepository),
	}
}

// CreateCarModelImage godoc
// @Summary Create a carModelImage
// @Description Create a carModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelImageRequest true "Create a carModelImage"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelImage, dto.ToCarModelImageResponse, h.carModelImageService.Create)
}

// UpdateCarModelImage godoc
// @Summary Update a carModelImage
// @Description Update a carModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a carModelImage"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelImage, dto.ToCarModelImageResponse, h.carModelImageService.Update)
}

// DeleteCarModelImage godoc
// @Summary Delete a carModelImage
// @Description Delete a carModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelImageService.Delete)
}

// GetCarModelImage godoc
// @Summary Get a carModelImage
// @Description Get a carModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelImageResponse} "carModelImage response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelImageResponse, h.carModelImageService.GetById)
}

// GetCarModelImages godoc
// @Summary Get CarModelImages
// @Description Get CarModelImages
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelImageResponse]} "CarModelImage response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelImageResponse, h.carModelImageService.GetByFilter)
}
