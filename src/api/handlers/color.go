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

type ColorHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	colorService *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	colorRepository := dependencies.GetColorRepository(cfg)

	return &ColorHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		colorService: services.NewColorService(cfg, colorRepository),
	}
}

// CreateColor godoc
// @Summary Create a color
// @Description Create a color
// @Tags Colors
// @Accept json
// @produces json
// @Param Request body dto.CreateColorRequest true "Create a color"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/colors/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateColor, dto.ToColorResponse, h.colorService.Create)
}

// UpdateColor godoc
// @Summary Update a color
// @Description Update a color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorRequest true "Update a color"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/colors/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateColor, dto.ToColorResponse, h.colorService.Update)
}

// DeleteColor godoc
// @Summary Delete a color
// @Description Delete a color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/colors/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c, h.colorService.Delete)
}

// GetColor godoc
// @Summary Get a color
// @Description Get a color
// @Tags Colors
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.ColorResponse} "color response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/colors/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToColorResponse, h.colorService.GetById)
}

// GetColors godoc
// @Summary Get Colors
// @Description Get Colors
// @Tags Colors
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.ColorResponse]} "Color response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/colors/get-by-filters [post]
// @Security AuthBearer
func (h *ColorHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToColorResponse, h.colorService.GetByFilter)
}
