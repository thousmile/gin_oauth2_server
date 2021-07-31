package handler

import (
	"gin_oauth2_server/domain"
	"gin_oauth2_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshHandler 刷新 token
func RefreshHandler(c *gin.Context) {
	token := c.GetHeader("RefreshToken")
	if oauth2Token, oAuth2Err := service.RefreshToken(token); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(oauth2Token))
	}
	return
}
