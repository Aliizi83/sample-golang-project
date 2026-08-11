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

type PersianYearHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	persianYearService *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	persianYearRepository := dependencies.GetPersianYearRepository(cfg)

	return &PersianYearHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		persianYearService: services.NewPersianYearService(cfg, persianYearRepository),
	}
}

// CreatePersianYear godoc
// @Summary Create a persianYear
// @Description Create a persianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param Request body dto.CreatePersianYearRequest true "Create a persianYear"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/persian-years/ [post]
// @Security AuthBearer
func (h *PersianYearHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreatePersianYear, dto.ToPersianYearResponse, h.persianYearService.Create)
}

// UpdatePersianYear godoc
// @Summary Update a persianYear
// @Description Update a persianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a persianYear"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PersianYearResponse} "PersianYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/persian-years/{id} [put]
// @Security AuthBearer
func (h *PersianYearHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdatePersianYear, dto.ToPersianYearResponse, h.persianYearService.Update)
}

// DeletePersianYear godoc
// @Summary Delete a persianYear
// @Description Delete a persianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/persian-years/{id} [delete]
// @Security AuthBearer
func (h *PersianYearHandler) Delete(c *gin.Context) {
	Delete(c, h.persianYearService.Delete)
}

// GetPersianYear godoc
// @Summary Get a persianYear
// @Description Get a persianYear
// @Tags PersianYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.PersianYearResponse} "persianYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/persian-years/{id} [get]
// @Security AuthBearer
func (h *PersianYearHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToPersianYearResponse, h.persianYearService.GetById)
}

// GetPersianYears godoc
// @Summary Get PersianYears
// @Description Get PersianYears
// @Tags PersianYears
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.PersianYearResponse]} "PersianYear response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/persian-years/get-by-filters [post]
// @Security AuthBearer
func (h *PersianYearHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToPersianYearResponse, h.persianYearService.GetByFilter)
}
