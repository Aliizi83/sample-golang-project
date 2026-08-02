package middlewares

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(getIP(c.Request.RemoteAddr))
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helpers.GenerateBaseResponseWithError(nil, false, int(helpers.OtpLimiterError), errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}
