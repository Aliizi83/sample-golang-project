package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelImage(r *gin.RouterGroup, cfg *config.Config) {
	carModelImageHandler := handlers.NewCarModelImageHandler(cfg)

	r.POST("/", carModelImageHandler.Create)
	r.PUT("/:id", carModelImageHandler.Update)
	r.DELETE("/:id", carModelImageHandler.Delete)
	r.GET("/:id", carModelImageHandler.GetById)
	r.POST("/get-by-filters", carModelImageHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-images", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelImage(router, cfg)
	})
}