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

type PropertyHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	propertyService *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	propertyRepository := dependencies.GetPropertyRepository(cfg)

	return &PropertyHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		propertyService: services.NewPropertyService(cfg, propertyRepository),
	}
}

// CreateProperty godoc
// @Summary Create a property
// @Description Create a property
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyRequest true "Create a property"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/properties/ [post]
// @Security AuthBearer
func (h *PropertyHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateProperty, dto.ToPropertyResponse, h.propertyService.Create)
}

// UpdateProperty godoc
// @Summary Update a property
// @Description Update a property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "Update a property"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [put]
// @Security AuthBearer
func (h *PropertyHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateProperty, dto.ToPropertyResponse, h.propertyService.Update)
}

// DeleteProperty godoc
// @Summary Delete a property
// @Description Delete a property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [delete]
// @Security AuthBearer
func (h *PropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.propertyService.Delete)
}

// GetProperty godoc
// @Summary Get a property
// @Description Get a property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.PropertyResponse} "property response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/properties/{id} [get]
// @Security AuthBearer
func (h *PropertyHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToPropertyResponse, h.propertyService.GetById)
}

// GetProperties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.PropertyResponse]} "Property response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/properties/get-by-filters [post]
// @Security AuthBearer
func (h *PropertyHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToPropertyResponse, h.propertyService.GetByFilter)
}
