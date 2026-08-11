package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModel(r *gin.RouterGroup, cfg *config.Config) {
	carModelHandler := handlers.NewCarModelHandler(cfg)

	r.POST("/", carModelHandler.Create)
	r.PUT("/:id", carModelHandler.Update)
	r.DELETE("/:id", carModelHandler.Delete)
	r.GET("/:id", carModelHandler.GetById)
	r.POST("/get-by-filters", carModelHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModel(router, cfg)
	})
}