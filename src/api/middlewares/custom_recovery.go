package middlewares

import (
	"net/http"

	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helpers.GenerateBaseResponseWithError(nil, false, int(helpers.CustomRecovery), err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helpers.GenerateBaseResponseWithAnyError(nil, false, helpers.CustomRecovery, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
