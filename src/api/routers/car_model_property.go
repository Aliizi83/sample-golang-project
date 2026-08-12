package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelProperty(r *gin.RouterGroup, cfg *config.Config) {
	carModelPropertyHandler := handlers.NewCarModelPropertyHandler(cfg)

	r.POST("/", carModelPropertyHandler.Create)
	r.PUT("/:id", carModelPropertyHandler.Update)
	r.DELETE("/:id", carModelPropertyHandler.Delete)
	r.GET("/:id", carModelPropertyHandler.GetById)
	r.POST("/get-by-filters", carModelPropertyHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelProperty(router, cfg)
	})
}