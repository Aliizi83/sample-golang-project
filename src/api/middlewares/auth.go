package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/pkg/service_errors"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMaps := map[string]interface{}{}
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" || len(token) < 2 {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMaps, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.GenerateBaseResponseWithError(
				nil, false, int(helpers.AuthError), err,
			))
			return
		}

		c.Set(constants.UserIdKey, claimMaps[constants.UserIdKey])
		c.Set(constants.FirstNameKey, claimMaps[constants.FirstNameKey])
		c.Set(constants.LastNameKey, claimMaps[constants.LastNameKey])
		c.Set(constants.UsernameKey, claimMaps[constants.UsernameKey])
		c.Set(constants.EmailKey, claimMaps[constants.EmailKey])
		c.Set(constants.MobileNumberKey, claimMaps[constants.MobileNumberKey])
		c.Set(constants.RolesKey, claimMaps[constants.RolesKey])
		c.Set(constants.ExpireTimeKey, claimMaps[constants.ExpireTimeKey])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, int(helpers.ForbiddenError)))
			return
		}

		rolesVal := c.Keys[constants.RolesKey]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, int(helpers.ForbiddenError)))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, helpers.GenerateBaseResponse(nil, false, int(helpers.ForbiddenError)))
	}
}
