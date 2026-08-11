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

type CarModelYearHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelYearService *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	carModelYearRepository := dependencies.GetCarModelYearRepository(cfg)

	return &CarModelYearHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelYearService: services.NewCarModelYearService(cfg, carModelYearRepository),
	}
}

// CreateCarModelYear godoc
// @Summary Create a carModelYear
// @Description Create a carModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelYearRequest true "Create a carModelYear"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelYear, dto.ToCarModelYearResponse, h.carModelYearService.Create)
}

// UpdateCarModelYear godoc
// @Summary Update a carModelYear
// @Description Update a carModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a carModelYear"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelYear, dto.ToCarModelYearResponse, h.carModelYearService.Update)
}

// DeleteCarModelYear godoc
// @Summary Delete a carModelYear
// @Description Delete a carModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelYearService.Delete)
}

// GetCarModelYear godoc
// @Summary Get a carModelYear
// @Description Get a carModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelYearResponse} "carModelYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelYearResponse, h.carModelYearService.GetById)
}

// GetCarModelYears godoc
// @Summary Get CarModelYears
// @Description Get CarModelYears
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelYearResponse]} "CarModelYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelYearResponse, h.carModelYearService.GetByFilter)
}
