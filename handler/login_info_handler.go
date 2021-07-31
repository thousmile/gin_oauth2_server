package handler

import (
	"gin_oauth2_server/config"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginInfoHandler 登录信息
func LoginInfoHandler(c *gin.Context) {
	if value, ok := c.Get(config.Config.OAuth2Server.TokenHeader); ok {
		tokenValue := value.(*domain.TokenValue)
		c.JSON(http.StatusOK, domain.Ok2(tokenValue))
		return
	}
	c.JSON(http.StatusOK, domain.Error2(exception.OAuth2Exception))
}

// LogoutHandler 退出登录
func LogoutHandler(c *gin.Context) {
	if value, ok := c.Get(config.Config.OAuth2Server.TokenHeader); ok {
		tokenValue := value.(*domain.TokenValue)
		// 退出登录
		service.Logout(tokenValue.LoginId)
		c.JSON(http.StatusOK, domain.Ok2(nil))
		return
	}
	c.JSON(http.StatusOK, domain.Error2(exception.OAuth2Exception))
}
