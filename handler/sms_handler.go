package handler

import (
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/params"
	"gin_oauth2_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

const mobilePattern = "^[1][3,4,5,7,8,9][0-9]{9}$"

// SendSmsHandler 发送 短信验证码
func SendSmsHandler(c *gin.Context) {
	var param params.SendSmsParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, domain.Error2(exception.RequestParamValidate))
		return
	}
	if len(param.ClientId) < 10 {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "client_id 格式错误！"))
		return
	}
	if ok, _ := regexp.MatchString(mobilePattern, param.Mobile); !ok {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "手机号格式不正确！"))
		return
	}
	if oAuth2Err := service.SendSms(param); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(nil))
	}
}

// SmsHandler 短信验证码 授权模式
func SmsHandler(c *gin.Context) {
	var param params.SmsModeParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, domain.Error2(exception.RequestParamValidate))
	}
	// 验证基础参数
	if err := param.Validator(); err != nil {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, err.Error()))
		return
	}
	if len(param.Code) < 4 {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "短信验证码格式错误！"))
		return
	}
	if ok, _ := regexp.MatchString(mobilePattern, param.Mobile); !ok {
		c.JSON(http.StatusOK, domain.Error3(exception.RequestParamValidate, "手机号格式不正确！"))
		return
	}
	if oauth2Token, oAuth2Err := service.SmsAuthorize(param); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(oauth2Token))
	}
}
