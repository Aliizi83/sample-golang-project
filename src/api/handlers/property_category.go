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

type PropertyCategoryHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	propertyCategoryService *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	propertyCategoryRepository := dependencies.GetPropertyCategoryRepository(cfg)

	return &PropertyCategoryHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		propertyCategoryService: services.NewPropertyCategoryService(cfg, propertyCategoryRepository),
	}
}

// CreatePropertyCategory godoc
// @Summary Create a propertyCategory
// @Description Create a propertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create a propertyCategory"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/ [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreatePropertyCategory, dto.ToPropertyCategoryResponse, h.propertyCategoryService.Create)
}

// UpdatePropertyCategory godoc
// @Summary Update a propertyCategory
// @Description Update a propertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a propertyCategory"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [put]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdatePropertyCategory, dto.ToPropertyCategoryResponse, h.propertyCategoryService.Update)
}

// DeletePropertyCategory godoc
// @Summary Delete a propertyCategory
// @Description Delete a propertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.propertyCategoryService.Delete)
}

// GetPropertyCategory godoc
// @Summary Get a propertyCategory
// @Description Get a propertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.PropertyCategoryResponse} "propertyCategory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/{id} [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToPropertyCategoryResponse, h.propertyCategoryService.GetById)
}

// GetPropertyCategories godoc
// @Summary Get PropertyCategories
// @Description Get PropertyCategories
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.PropertyCategoryResponse]} "PropertyCategory response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/get-by-filters [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToPropertyCategoryResponse, h.propertyCategoryService.GetByFilter)
}
