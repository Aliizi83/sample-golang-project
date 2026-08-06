package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func files(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filters", h.GetByFilters)
}

func init() {
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		files(router, cfg)
	})
}
