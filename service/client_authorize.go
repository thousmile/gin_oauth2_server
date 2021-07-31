package service

import (
	"gin_oauth2_server/config"
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/mapper"
	"gin_oauth2_server/params"
	"gin_oauth2_server/util"
	"github.com/jameskeane/bcrypt"
	"time"
)

var conf = config.Config.OAuth2Server

// GetClient 根据客户端ID,查询客户端
func GetClient(clientId string) (*domain.ClientDetails, *exception.OAuth2Error) {
	var myMapper = mapper.BaseMapper
	var client domain.ClientDetails
	myMapper.Where("client_id = ? ", clientId).First(&client)
	if len(client.ClientId) < 1 || len(client.Secret) < 1 {
		return nil, exception.ClientInvalid
	}
	return &client, nil
}

// ValidateClient 校验客户端
func ValidateClient(clientId, clientSecret, grantType string) (*domain.ClientDetails, *exception.OAuth2Error) {
	client, err := GetClient(clientId)
	if err != nil {
		return nil, err
	}
	// 客户端被禁用
	if client.Status == constant.Disable {
		return nil, exception.ClientDisable
	}

	// 判断客户端是否包含此 授权模式
	if !client.ContainsGrantType(grantType) {
		return nil, exception.AuthorizationGrantType
	}

	// 校验 密钥
	if !bcrypt.Match(clientSecret, client.Secret) {
		return nil, exception.ClientSecretError
	}
	return client, nil
}

// BuildToken 构建 OAuth2Token
func BuildToken(loginId, scope string) (*domain.OAuth2Token, *exception.OAuth2Error) {
	accessToken, err := util.CreateAccessToken(loginId)
	if err != nil {
		return nil, &exception.OAuth2Error{Status: 500, Message: err.Error()}
	}
	refreshToken, _ := util.CreateRefreshToken(loginId)
	return &domain.OAuth2Token{
		AccessToken:  accessToken,
		TokenType:    conf.TokenType,
		RefreshToken: refreshToken,
		Scope:        scope,
		ExpiresIn:    conf.TokenExpired,
	}, nil
}

// ClientAuthorize 客户端授权模式
func ClientAuthorize(param params.ClientModeParam) (*domain.OAuth2Token, *exception.OAuth2Error) {
	// 校验客户端
	client, auth2Error := ValidateClient(
		param.ClientId,
		param.ClientSecret,
		constant.ClientCredentials,
	)

	if auth2Error != nil {
		return nil, auth2Error
	}

	// 唯一的 登录ID
	loginId := util.CreateLoginId()

	tokenValue := &domain.TokenValue{
		LoginId:   loginId,
		LoginTime: time.Now().Unix(),
		Client:    client,
		GrantType: constant.ClientCredentials,
		User:      nil,
	}

	SetAccessToken(loginId, tokenValue)

	SetRefreshToken(loginId, tokenValue)

	return BuildToken(loginId, client.Scope)
}
