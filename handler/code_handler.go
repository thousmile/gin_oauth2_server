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

// GetCodeHandler 授权码模式，1.第三方应用，请求获取用户信息。在此校验客户端是否正确
func GetCodeHandler(c *gin.Context) {
	var param params.GetCodeModeParam
	if err := c.ShouldBind(&param); err != nil {
		c.HTML(http.StatusOK, "paramError.tmpl", domain.Error3(
			exception.OAuth2Exception, err.Error()),
		)
		return
	}

	// 客户端授权模式 GrantType 必须 code
	if !strings.EqualFold(constant.Code, param.ResponseType) {
		c.HTML(http.StatusOK, "paramError.tmpl", domain.Error3(
			exception.RequestParamValidate,
			"授权码模式 response_type 必须 code"),
		)
		return
	}
	param.ClientId = strings.Trim(param.ClientId, " ")
	if codeId, oAuth2Err := service.CodeValidateClient(&param); oAuth2Err != nil {
		c.HTML(http.StatusOK, "paramError.tmpl", domain.Error2(oAuth2Err))
	} else {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"codeId": *codeId,
		})
	}
	return
}

// LoginHandler 授权码模式，第二步，用户登录
func LoginHandler(c *gin.Context) {
	var param params.UserLoginParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, domain.Error3(exception.OAuth2Exception, err.Error()))
		return
	}

	// 必须携带 code id
	if len(param.CodeId) < 5 {
		c.JSON(http.StatusOK, domain.Error3(exception.CodeInvalid, "授权码模式 code 不存在！"))
		return
	}

	// 验证登录信息，生成回调 url
	if redirectUri, oAuth2Err := service.UserLogin(param); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(redirectUri))
	}

	return
}

// GetAccessToken 授权码模式，第三步，第三方应用通过 code 来换取 access_token
func GetAccessToken(c *gin.Context) {
	var param params.CodeModeParam
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, domain.Error3(exception.OAuth2Exception, err.Error()))
		return
	}
	// 必须携带 code
	if len(param.Code) < 5 {
		c.JSON(http.StatusOK, domain.Error2(exception.CodeInvalid))
		return
	}
	// 授权码模式 GrantType 必须 authorization_code
	if !strings.EqualFold(constant.AuthorizationCode, param.GrantType) {
		c.JSON(http.StatusOK, domain.Error3(
			exception.RequestParamValidate,
			"授权码模式 GrantType 必须 authorization_code"),
		)
		return
	}
	// 验证 code 信息，生成 token
	if oauth2Token, oAuth2Err := service.CodeAuthorize(param); oAuth2Err != nil {
		c.JSON(http.StatusOK, domain.Error2(oAuth2Err))
	} else {
		c.JSON(http.StatusOK, domain.Ok2(oauth2Token))
	}
	return
}
