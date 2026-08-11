package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelColor(r *gin.RouterGroup, cfg *config.Config) {
	carModelColorHandler := handlers.NewCarModelColorHandler(cfg)

	r.POST("/", carModelColorHandler.Create)
	r.PUT("/:id", carModelColorHandler.Update)
	r.DELETE("/:id", carModelColorHandler.Delete)
	r.GET("/:id", carModelColorHandler.GetById)
	r.POST("/get-by-filters", carModelColorHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelColor(router, cfg)
	})
}