package middleware

import (
	"gin_oauth2_server/config"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var conf = config.Config.OAuth2Server

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader(conf.TokenHeader)
		tokenValue, auth2Error := service.Validate(bearerToken)
		if auth2Error != nil {
			c.Abort()
			c.JSON(http.StatusOK, domain.Error2(auth2Error))
			return
		}
		c.Set(conf.TokenHeader, tokenValue)
		c.Next()
	}
}
