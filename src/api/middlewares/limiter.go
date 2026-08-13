package middlewares

import (
	"net/http"

	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helpers.GenerateBaseResponseWithError(nil, false, int(helpers.LimiterError), err))
			return
		} else {
			c.Next()
		}
	}
}
