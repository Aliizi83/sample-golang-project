package routers

import (
	"github.com/Aliizi83/sample-golang-project/src/api/handlers"
	"github.com/Aliizi83/sample-golang-project/src/api/middlewares"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/gin-gonic/gin"
)

func user(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUserHandler(cfg)

	router.POST("/send-otp", h.SendOtp, middlewares.OtpLimiter(cfg))
	router.POST("/register-by-username", h.RegisterUserByUsername)
	router.POST("/register-login-by-mobile", h.RegisterLoginByMobile)
	router.POST("/login-by-username-password", h.LoginByUsername)
}

func init() {
	RegisterRoute(func(v1 *gin.RouterGroup, cfg *config.Config) {
		router := v1.Group("/users", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		user(router, cfg)
	})
}
