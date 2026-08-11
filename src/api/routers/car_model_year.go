package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelYear(r *gin.RouterGroup, cfg *config.Config) {
	carModelYearHandler := handlers.NewCarModelYearHandler(cfg)

	r.POST("/", carModelYearHandler.Create)
	r.PUT("/:id", carModelYearHandler.Update)
	r.DELETE("/:id", carModelYearHandler.Delete)
	r.GET("/:id", carModelYearHandler.GetById)
	r.POST("/get-by-filters", carModelYearHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelYear(router, cfg)
	})
}