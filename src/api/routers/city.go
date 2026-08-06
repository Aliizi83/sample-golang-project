package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func cities(r *gin.RouterGroup, cfg *config.Config) {
	cityHandler := handlers.NewCityHandler(cfg)

	r.POST("/", cityHandler.Create)
	r.PUT("/:id", cityHandler.Update)
	r.DELETE("/:id", cityHandler.Delete)
	r.GET("/:id", cityHandler.GetById)
	r.POST("/get-by-filters", cityHandler.GetByFilters)
}

func init() {
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities(router, cfg)
	})
}
