package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func carModelComment(r *gin.RouterGroup, cfg *config.Config) {
	carModelCommentHandler := handlers.NewCarModelCommentHandler(cfg)

	r.POST("/", carModelCommentHandler.Create)
	r.PUT("/:id", carModelCommentHandler.Update)
	r.DELETE("/:id", carModelCommentHandler.Delete)
	r.GET("/:id", carModelCommentHandler.GetById)
	r.POST("/get-by-filters", carModelCommentHandler.GetByFilters)
}

func init(){
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/car-model-comments", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelComment(router, cfg)
	})
}