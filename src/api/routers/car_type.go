package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carType(r *gin.RouterGroup, cfg *config.Config) {
	carTypeHandler := handlers.NewCarTypeHandler(cfg)

	r.POST("/", carTypeHandler.Create)
	r.PUT("/:id", carTypeHandler.Update)
	r.DELETE("/:id", carTypeHandler.Delete)
	r.GET("/:id", carTypeHandler.GetById)
	r.POST("/get-by-filters", carTypeHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carType(router, cfg)
	})
}