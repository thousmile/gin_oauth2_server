package service

import (
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/mapper"
	"gin_oauth2_server/params"
	"gin_oauth2_server/util"
	"time"
)

// GetUser 根据 Username Mobile Email 获取用户信息
func GetUser(account string) (*domain.UserInfo, *exception.OAuth2Error) {
	var myMapper = mapper.BaseMapper
	var user domain.UserInfo
	// 先根据 Username 查询用户
	myMapper.Where("username = ?", account).First(&user)
	if len(user.Username) < 1 || len(user.UserId) < 1 {
		// 再根据 Mobile 查询用户
		myMapper.Where("mobile = ?", account).First(&user)
		if len(user.Username) < 1 || len(user.UserId) < 1 {
			// 再根据 Email 查询用户
			myMapper.Where("email = ?", account).First(&user)
		}
	}
	if len(user.Username) < 1 || len(user.UserId) < 1 {
		return nil, exception.UserInvalid
	}
	return &user, nil
}

func ValidateUser(account, password string) (*domain.UserInfo, *exception.OAuth2Error) {
	dbUser, err := GetUser(account)
	if err != nil {
		return nil, err
	}
	// 用户被锁定
	if dbUser.Status == constant.Locking {
		return nil, exception.UserLocking
	}
	// 用户被禁用
	if dbUser.Status == constant.Disable {
		return nil, exception.UserDisable
	}
	// 校验密码
	if !util.PasswordMatches(password, dbUser.Password) {
		return nil, exception.UserPasswordError
	}
	return dbUser, nil
}

// PasswordAuthorize  密码授权模式
func PasswordAuthorize(param params.PasswordModeParam) (*domain.OAuth2Token, *exception.OAuth2Error) {
	var auth2Error *exception.OAuth2Error
	var client *domain.ClientDetails
	// 校验 客户端
	if client, auth2Error = ValidateClient(param.ClientId, param.ClientSecret, constant.Password); auth2Error != nil {
		return nil, auth2Error
	}
	// 校验 用户信息
	var user *domain.UserInfo
	if user, auth2Error = ValidateUser(param.Username, param.Password); auth2Error != nil {
		return nil, auth2Error
	}
	// 唯一的 登录ID
	loginId := util.CreateLoginId()
	tokenValue := &domain.TokenValue{
		LoginId:   loginId,
		LoginTime: time.Now().Unix(),
		Client:    client,
		GrantType: constant.Password,
		User:      user,
	}

	SetAccessToken(loginId, tokenValue)

	SetRefreshToken(loginId, tokenValue)

	return BuildToken(loginId, client.Scope)
}
