package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func property(r *gin.RouterGroup, cfg *config.Config) {
	propertyHandler := handlers.NewPropertyHandler(cfg)

	r.POST("/", propertyHandler.Create)
	r.PUT("/:id", propertyHandler.Update)
	r.DELETE("/:id", propertyHandler.Delete)
	r.GET("/:id", propertyHandler.GetById)
	r.POST("/get-by-filters", propertyHandler.GetByFilters)
}

func init() {
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		property(router, cfg)
	})
}
