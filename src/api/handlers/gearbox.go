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

type GearboxHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	gearboxService *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	gearboxRepository := dependencies.GetGearboxRepository(cfg)

	return &GearboxHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		gearboxService: services.NewGearboxService(cfg, gearboxRepository),
	}
}

// CreateGearbox godoc
// @Summary Create a gearbox
// @Description Create a gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.CreateGearboxRequest true "Create a gearbox"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateGearbox, dto.ToGearboxResponse, h.gearboxService.Create)
}

// UpdateGearbox godoc
// @Summary Update a gearbox
// @Description Update a gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a gearbox"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateGearbox, dto.ToGearboxResponse, h.gearboxService.Update)
}

// DeleteGearbox godoc
// @Summary Delete a gearbox
// @Description Delete a gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.gearboxService.Delete)
}

// GetGearbox godoc
// @Summary Get a gearbox
// @Description Get a gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.GearboxResponse} "gearbox response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToGearboxResponse, h.gearboxService.GetById)
}

// GetGearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.GearboxResponse]} "Gearbox response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/get-by-filters [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToGearboxResponse, h.gearboxService.GetByFilter)
}
