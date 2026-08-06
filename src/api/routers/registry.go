package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

type RouteRegisterFunc func(v1 *gin.RouterGroup, cfg *config.Config)

var RegisteredRoutes []RouteRegisterFunc

func RegisterRoute(fn RouteRegisterFunc) {
	RegisteredRoutes = append(RegisteredRoutes, fn)
}
