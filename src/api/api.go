package api

import (
	"fmt"

	"github.com/Aliizi83/sample-golang-project/src/api/routers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {

	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

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

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}
