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

type CarModelCommentHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	carModelCommentService *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	carModelCommentRepository := dependencies.GetCarModelCommentRepository(cfg)

	return &CarModelCommentHandler{
		cfg:         cfg,
		logger:      logging.NewLogger(cfg),
		carModelCommentService: services.NewCarModelCommentService(cfg, carModelCommentRepository),
	}
}

// CreateCarModelComment godoc
// @Summary Create a carModelComment
// @Description Create a carModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelCommentRequest true "Create a carModelComment"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelComment, dto.ToCarModelCommentResponse, h.carModelCommentService.Create)
}

// UpdateCarModelComment godoc
// @Summary Update a carModelComment
// @Description Update a carModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a carModelComment"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelComment, dto.ToCarModelCommentResponse, h.carModelCommentService.Update)
}

// DeleteCarModelComment godoc
// @Summary Delete a carModelComment
// @Description Delete a carModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.carModelCommentService.Delete)
}

// GetCarModelComment godoc
// @Summary Get a carModelComment
// @Description Get a carModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CarModelCommentResponse} "carModelComment response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelCommentResponse, h.carModelCommentService.GetById)
}

// GetCarModelComments godoc
// @Summary Get CarModelComments
// @Description Get CarModelComments
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.CarModelCommentResponse]} "CarModelComment response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/get-by-filters [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelCommentResponse, h.carModelCommentService.GetByFilter)
}
