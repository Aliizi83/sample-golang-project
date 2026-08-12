package api

import (
	"fmt"

	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/api/routers"
	"github.com/Aliizi83/sample-golang-project/src/api/validations"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/docs"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/metrics"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	goSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {

	r := gin.New()
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.Prometheus())
	RegisterPrometheus()
	RegisterValidators()
	RegisterRouters(r, cfg)
	RegisterSwagger(r, cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator); err != nil {
			logger.Error(err, logging.Validation, logging.Startup, err.Error(), nil)
		}
		if err := val.RegisterValidation("password", validations.PasswordValidator); err != nil {
			logger.Error(err, logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}

func RegisterRouters(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			health := v1.Group("/health")
			routers.Health(health)

			for _, routerFunction := range routers.RegisteredRoutes {
				routerFunction(v1, cfg)
			}

			r.Static("/static", "./uploads")
			r.GET("/metrics", gin.WrapH(promhttp.Handler()))
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

func RegisterPrometheus() {
	if err := prometheus.Register(metrics.DbCall); err != nil {
		logger.Error(err, logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	if err := prometheus.Register(metrics.HttpDuration); err != nil {
		logger.Error(err, logging.Prometheus, logging.Startup, err.Error(), nil)
	}

}
