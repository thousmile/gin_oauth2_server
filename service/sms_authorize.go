package service

import (
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/log"
	"gin_oauth2_server/mapper"
	"gin_oauth2_server/params"
	"gin_oauth2_server/util"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetUserMobile(mobile string) (*domain.UserInfo, *exception.OAuth2Error) {
	var user domain.UserInfo
	// 根据 Mobile 查询用户
	mapper.BaseMapper.Where("mobile = ?", mobile).First(&user)
	if len(user.Username) < 1 || len(user.UserId) < 1 {
		return nil, exception.UserMobileInvalid
	}
	return &user, nil
}

// 发送短信验证码
func sendAliyunSms(mobile, code string) {
	log.Logger.Debugf("发送验证码 [ %s ] 到 %s", code, mobile)
}

func SendSms(param params.SendSmsParam) *exception.OAuth2Error {
	var auth2Error *exception.OAuth2Error
	var client *domain.ClientDetails
	if client, auth2Error = GetClient(param.ClientId); auth2Error != nil {
		return auth2Error
	}
	// 校验 用户信息
	var user *domain.UserInfo
	if user, auth2Error = GetUserMobile(param.Mobile); auth2Error != nil {
		return auth2Error
	}
	// 短信验证码 code
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	code := strconv.Itoa(num)

	// 发送短信验证码
	sendAliyunSms(user.Mobile, code)

	expired := time.Duration(conf.SmsCodeExpired) * time.Second

	// 将 短信验证码保存起来。
	SetString(constant.SmsKey+user.Mobile, code, expired)

	// 将客户端ID，也保存起来
	SetString(constant.ClientKey+user.Mobile, client.ClientId, expired)

	return nil
}

func SmsAuthorize(param params.SmsModeParam) (*domain.OAuth2Token, *exception.OAuth2Error) {
	// 校验发送短信的客户端，和本次请求的客户端是否一致
	oldClientId := GetString(constant.ClientKey + param.Mobile)
	if !strings.EqualFold(param.ClientId, oldClientId) {
		return nil, exception.ClientIsDifferent
	}

	// 校验 短信验证码
	code := GetString(constant.SmsKey + param.Mobile)
	if !strings.EqualFold(code, param.Code) {
		return nil, exception.SmsCodeError
	}

	var auth2Error *exception.OAuth2Error
	var client *domain.ClientDetails
	// 校验 客户端
	if client, auth2Error = ValidateClient(param.ClientId, param.ClientSecret, constant.Password); auth2Error != nil {
		return nil, auth2Error
	}

	// 校验 用户信息
	var user *domain.UserInfo
	if user, auth2Error = GetUserMobile(param.Mobile); auth2Error != nil {
		return nil, auth2Error
	}

	// 唯一的 登录ID
	loginId := util.CreateLoginId()
	tokenValue := &domain.TokenValue{
		LoginId:   loginId,
		LoginTime: time.Now().Unix(),
		Client:    client,
		GrantType: constant.Sms,
		User:      user,
	}

	// 将 短信验证码 移除内存
	Remove(constant.SmsKey + user.Mobile)

	// 将 客户端ID 移除内存
	Remove(constant.ClientKey + user.Mobile)

	SetAccessToken(loginId, tokenValue)

	SetRefreshToken(loginId, tokenValue)

	return BuildToken(loginId, client.Scope)
}
