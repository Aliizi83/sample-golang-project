package handlers

import (
	_ "github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} helpers.BaseHttpResponse "Success"
// @Failure 400 {object} helpers.BaseHttpResponse "Failed"
// @Router /v1/health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, "Working")
}
