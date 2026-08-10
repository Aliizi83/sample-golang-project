package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func color(r *gin.RouterGroup, cfg *config.Config) {
	colorHandler := handlers.NewColorHandler(cfg)

	r.POST("/", colorHandler.Create)
	r.PUT("/:id", colorHandler.Update)
	r.DELETE("/:id", colorHandler.Delete)
	r.GET("/:id", colorHandler.GetById)
	r.POST("/get-by-filters", colorHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		color(router, cfg)
	})
}