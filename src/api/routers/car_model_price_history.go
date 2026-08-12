package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelPriceHistory(r *gin.RouterGroup, cfg *config.Config) {
	carModelPriceHistoryHandler := handlers.NewCarModelPriceHistoryHandler(cfg)

	r.POST("/", carModelPriceHistoryHandler.Create)
	r.PUT("/:id", carModelPriceHistoryHandler.Update)
	r.DELETE("/:id", carModelPriceHistoryHandler.Delete)
	r.GET("/:id", carModelPriceHistoryHandler.GetById)
	r.POST("/get-by-filters", carModelPriceHistoryHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-price-histories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelPriceHistory(router, cfg)
	})
}