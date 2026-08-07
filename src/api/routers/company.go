package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func company(r *gin.RouterGroup, cfg *config.Config) {
	companyHandler := handlers.NewCompanyHandler(cfg)

	r.POST("/", companyHandler.Create)
	r.PUT("/:id", companyHandler.Update)
	r.DELETE("/:id", companyHandler.Delete)
	r.GET("/:id", companyHandler.GetById)
	r.POST("/get-by-filters", companyHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		company(router, cfg)
	})
}