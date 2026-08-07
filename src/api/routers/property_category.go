package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func propertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	propertyCategoryHandler := handlers.NewPropertyCategoryHandler(cfg)

	r.POST("/", propertyCategoryHandler.Create)
	r.PUT("/:id", propertyCategoryHandler.Update)
	r.DELETE("/:id", propertyCategoryHandler.Delete)
	r.GET("/:id", propertyCategoryHandler.GetById)
	r.POST("/get-by-filters", propertyCategoryHandler.GetByFilters)
}

func init() {
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/property-categories")
		propertyCategory(router, cfg)
	})
}
