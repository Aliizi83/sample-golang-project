package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func Cities(r *gin.RouterGroup, cfg *config.Config) {
	cityHandler := handlers.NewCityHandler(cfg)

	r.POST("/", cityHandler.Create)
	r.PUT("/:id", cityHandler.Update)
	r.DELETE("/:id", cityHandler.Delete)
	r.GET("/:id", cityHandler.GetById)
	r.GET("/get-by-filters", cityHandler.GetByFilters)
}
