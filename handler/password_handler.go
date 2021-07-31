package handler

import (
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/params"
	"gin_oauth2_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// PasswordHandler 密码授权模式
func PasswordHandler(c *gin.Context) {
	var param params.PasswordModeParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, domain.Error1(500, err.Error()))
		return
	}
	// 验证基础参数
	if err := param.Validator(); err != nil {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, err.Error()))
		return
	}
	if len(param.Username) < 5 {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "用户名必须填写，并且长度大于5位！"))
		return
	}
	if len(param.Password) < 5 {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "密码必须填写，并且长度大于5位！"))
		return
	}
	// 客户端授权模式 GrantType 必须 client_credentials
	if !strings.EqualFold(constant.Password, param.GrantType) {
		c.JSON(http.StatusOK, domain.Error3(
			exception.RequestParamValidate,
			"密码授权模式 GrantType 必须 password"),
		)
		return
	}
	if oauth2Token, oAuth2Err := service.PasswordAuthorize(param); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(oauth2Token))
	}
	return
}
