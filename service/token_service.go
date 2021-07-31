package service

import (
	"fmt"
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/util"
	"strings"
)

// token 校验
func Validate(bearerToken string) (*domain.TokenValue, *exception.OAuth2Error) {
	loginId, auth2Error := util.GetIdFromToken(bearerToken)
	if auth2Error != nil {
		return nil, auth2Error
	}
	tokenValue, _ := GetAccessToken(loginId)
	// 判断 token 是否存在
	if tokenValue == nil || tokenValue.Client == nil {
		// 判断是否启用单点登录
		if conf.Sso {
			forcedOfflineKey := constant.ForcedOfflineKey + loginId
			// 判断此用户，是不是被挤下线
			offlineTime := GetString(forcedOfflineKey)
			if len(offlineTime) > 1 {
				// 删除 被挤下线 的消息提示
				Remove(forcedOfflineKey)
				err := exception.OAuth2Exception
				err.Message = fmt.Sprintf("您的账号在[ %s ]被其他用户拥下线了！", offlineTime)
				return nil, err
			}
		}
		err := exception.OAuth2Exception
		err.Message = "当前登录用户不存在！"
		return nil, err
	}

	return tokenValue, nil
}

// Logout 退出登录
func Logout(loginId string) *exception.OAuth2Error {
	tokenValue, auth2Error := GetAccessToken(loginId)
	if auth2Error != nil {
		return auth2Error
	}
	// 判断客户端类型
	if !strings.EqualFold(constant.ClientCredentials, tokenValue.GrantType) {
		Username := tokenValue.User.Username
		// 移出 在线用户
		Remove(constant.OnlineUserKey + Username)
		// 移出 强制下线
		Remove(constant.ForcedOfflineKey + Username)
	}

	// 移出 token
	Remove(constant.LoginTokenKey + loginId)

	// 移出刷新token
	Remove(constant.RefreshTokenKey + loginId)
	return nil
}

// RefreshToken  刷新 token
func RefreshToken(bearerToken string) (*domain.OAuth2Token, *exception.OAuth2Error) {
	oldLoginId, auth2Error := util.GetIdFromToken(bearerToken)
	if auth2Error != nil {
		return nil, auth2Error
	}
	tokenValue, auth2Error := GetRefreshToken(oldLoginId)
	if tokenValue == nil || len(tokenValue.LoginId) < 5 {
		return nil, auth2Error
	}

	// 创建一个新的 loginId
	newLoginId := util.CreateLoginId()

	tokenValue.LoginId = newLoginId
	SetAccessToken(newLoginId, tokenValue)
	SetRefreshToken(newLoginId, tokenValue)

	Logout(oldLoginId)

	// 刷新完成 token 后，会把之前的 oldLoginId 强制下线！
	Remove(constant.ForcedOfflineKey + oldLoginId)

	return BuildToken(newLoginId, tokenValue.Client.Scope)
}
