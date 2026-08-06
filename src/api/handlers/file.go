package handlers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"

	_ "github.com/Aliizi83/sample-golang-project/src/api/helpers"
	_ "github.com/Aliizi83/sample-golang-project/src/domain/filters"
)

type FileHandler struct {
	fileService *services.FileService
	cfg         *config.Config
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		fileService: services.NewFileService(cfg),
		cfg:         cfg,
	}
}

// CreateFile godoc
// @Summary Create a file
// @Description Create a file
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param description formData string true "File description"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/files/ [post]
// @Security AuthBearer
func (h *FileHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateFile, dto.ToFileResponse, h.fileService.Create)
}

// UpdateFile godoc
// @Summary Update a file
// @Description Update a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateFileRequest true "Update a file"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [put]
// @Security AuthBearer
func (h *FileHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateFile, dto.ToFileResponse, h.fileService.Update)
}

// DeleteFile godoc
// @Summary Delete a file
// @Description Delete a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) Delete(c *gin.Context) {
	Delete(c, h.fileService.Delete)
}

// GetFile godoc
// @Summary Get a file
// @Description Get a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToFileResponse, h.fileService.GetById)
}

// GetFiles godoc
// @Summary Get Files
// @Description Get Files
// @Tags Files
// @Accept json
// @produces json
// @Param Request body filters.PaginationInputWithFilter true "Request"
// @Success 200 {object} helpers.BaseHttpResponse{result=filters.PagedList[dto.FileResponse]} "File response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/files/get-by-filters [post]
// @Security AuthBearer
func (h *FileHandler) GetByFilters(c *gin.Context) {
	GetByFilter(c, dto.ToFileResponse, h.fileService.GetByFilter)
}
