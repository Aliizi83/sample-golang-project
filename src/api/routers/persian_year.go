package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func persianYear(r *gin.RouterGroup, cfg *config.Config) {
	persianYearHandler := handlers.NewPersianYearHandler(cfg)

	r.POST("/", persianYearHandler.Create)
	r.PUT("/:id", persianYearHandler.Update)
	r.DELETE("/:id", persianYearHandler.Delete)
	r.GET("/:id", persianYearHandler.GetById)
	r.POST("/get-by-filters", persianYearHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/persian-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		persianYear(router, cfg)
	})
}