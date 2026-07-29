package api

import (
	"fmt"

	"github.com/Aliizi83/sample-golang-project/src/api/routers"
	"github.com/Aliizi83/sample-golang-project/src/api/validations"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	goSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterRouters(r)
	RegisterSwagger(r, cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator)
		val.RegisterValidation("password", validations.PasswordValidator)
	}
}

func RegisterRouters(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			health := v1.Group("/health")
			routers.Health(health)
		}

		v2 := api.Group("/v2")
		{
			health := v2.Group("/health")
			routers.Health(health)

		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Web api documentation"
	docs.SwaggerInfo.Description = "Web api documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Domain, cfg.Server.InternalPort)
	r.GET("/swagger/*any", goSwagger.WrapHandler(swaggerFiles.Handler))
}
