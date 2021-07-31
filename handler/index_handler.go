package handler

import (
	conf "gin_oauth2_server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandler 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": conf.Config.AppName,
	})
	return
}
