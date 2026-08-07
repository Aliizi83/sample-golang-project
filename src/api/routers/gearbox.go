package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func gearbox(r *gin.RouterGroup, cfg *config.Config) {
	gearboxHandler := handlers.NewGearboxHandler(cfg)

	r.POST("/", gearboxHandler.Create)
	r.PUT("/:id", gearboxHandler.Update)
	r.DELETE("/:id", gearboxHandler.Delete)
	r.GET("/:id", gearboxHandler.GetById)
	r.POST("/get-by-filters", gearboxHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		gearbox(router, cfg)
	})
}