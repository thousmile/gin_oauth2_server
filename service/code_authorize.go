package service

import (
	"fmt"
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/params"
	"gin_oauth2_server/util"
	"net/url"
	"strings"
	"time"
)

// CodeValidateClient 校验客户端
func CodeValidateClient(param *params.GetCodeModeParam) (*string, *exception.OAuth2Error) {
	client, err := GetClient(param.ClientId)
	if err != nil {
		return nil, err
	}
	// 客户端被禁用
	if client.Status == constant.Disable {
		return nil, exception.ClientDisable
	}
	// 判断客户端是否包含此 授权模式
	if !client.ContainsGrantType(param.ResponseType) {
		return nil, exception.AuthorizationGrantType
	}
	// 判断 域名是否一致
	paramUrl, err1 := url.Parse(param.RedirectUri)
	if err1 != nil {
		return nil, exception.DomainNameIllegal
	}
	if !strings.EqualFold(client.DomainName, paramUrl.Host) {
		return nil, exception.DomainNameIllegal
	}

	codeId := util.CreateLoginId()
	// 将信息传入的信息保存在 缓存中
	SetCodeModeParam(codeId, param)

	return &codeId, nil
}

// 用户登录
func UserLogin(param params.UserLoginParam) (*string, *exception.OAuth2Error) {
	codeParam, auth2Error := GetCodeModeParam(param.CodeId)
	if auth2Error != nil {
		return nil, auth2Error
	}
	// 校验 用户信息
	var user *domain.UserInfo
	if user, auth2Error = ValidateUser(param.Username, param.Password); auth2Error != nil {
		return nil, auth2Error
	}

	// 将登录的用户信息，缓存起来
	SetLoginUserInfo(param.CodeId, user)

	// 构造 Redirect Uri
	/**
	 * 判断url是否已经携带参数了
	 * 如：https://xaaef.com/authorize/callback?pid=2135543
	 * 就只能在后面继续加 &
	 * 如：https://xaaef.com/authorize/callback
	 * 就需要加 ?
	 * */
	var prefix = "?"
	if strings.Contains(codeParam.RedirectUri, prefix) {
		prefix = "&"
	}
	redirectUri := fmt.Sprintf("%s%scode=%s&state=%s",
		codeParam.RedirectUri, prefix, param.CodeId, codeParam.State,
	)
	return &redirectUri, nil
}

// CodeAuthorize 授权码模式
func CodeAuthorize(param params.CodeModeParam) (*domain.OAuth2Token, *exception.OAuth2Error) {
	loginUser, auth2Error := GetLoginUserInfo(param.Code)
	if auth2Error != nil {
		return nil, auth2Error
	}

	// 校验两次客户端是否一致
	modeParam, auth2Error := GetCodeModeParam(param.Code)
	if auth2Error != nil || !strings.EqualFold(param.ClientId, modeParam.ClientId) {
		return nil, exception.ClientIsDifferent
	}

	// 校验客户端 密钥是否正确
	client, auth2Error := ValidateClient(
		param.ClientId,
		param.ClientSecret,
		constant.Code,
	)

	if auth2Error != nil {
		return nil, auth2Error
	}
	// param.Code
	// 唯一的 登录ID
	loginId := util.CreateLoginId()

	tokenValue := &domain.TokenValue{
		LoginId:   loginId,
		LoginTime: time.Now().Unix(),
		Client:    client,
		GrantType: constant.Code,
		User:      loginUser,
	}

	SetAccessToken(loginId, tokenValue)

	SetRefreshToken(loginId, tokenValue)

	// 移除 登录的用户
	Remove(constant.CodeLoginUserKey + param.Code)

	// 移除 授权的 客户端 参数信息
	Remove(constant.CodeModeKey + param.Code)

	return BuildToken(loginId, client.Scope)
}
