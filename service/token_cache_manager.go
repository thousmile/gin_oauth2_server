package service

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_oauth2_server/config"
	"gin_oauth2_server/constant"
	"gin_oauth2_server/domain"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/log"
	"gin_oauth2_server/params"
	"gin_oauth2_server/util"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"net/http"
	"os"
	"strings"
	"time"
)

var goCache *cache.Cache

var redisCache *redis.Client

var ctx = context.Background()

var cacheType = strings.ToLower(config.Config.CacheType)

func init() {
	switch cacheType {
	case "redis":
		log.Logger.Debug("redis cache...")
		rc := config.Config.Redis
		redisCache = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
			Password: rc.Password, // no password set
			DB:       rc.Db,       // use default DB
		})
		_, err := redisCache.Ping(ctx).Result()
		if err != nil {
			panic(fmt.Sprintf("%v", err))
			os.Exit(1)
		}
		break
	default:
		log.Logger.Debug("go-cache...")
		goCache = cache.New(5*time.Minute, 10*time.Minute)
		break
	}
}

// SetAccessToken 设置 AccessToken
func SetAccessToken(loginId string, token *domain.TokenValue) {
	jsonStr, err := json.Marshal(token)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	SetString(constant.LoginTokenKey+loginId, string(jsonStr), time.Duration(conf.TokenExpired)*time.Second)

	// 如果是 客户端模式，那么 就没有 单点登录
	if strings.EqualFold(constant.ClientCredentials, token.GrantType) {
		return
	}

	// 其他授权模式，判断是否开启 单点登录
	if conf.Sso {
		user := token.User
		// 在线用户，
		onlineUserKey := constant.OnlineUserKey + user.Username

		// 获取使用 此用户登录的 上个用户的 loginId
		oldLoginId := GetString(onlineUserKey)
		if len(oldLoginId) > 0 {
			// 移除 之前登录的 用户 token 以及 refresh_token
			Remove(constant.LoginTokenKey + oldLoginId)
			Remove(constant.RefreshTokenKey + oldLoginId)
			// 移除 之前登录的在线用户
			Remove(onlineUserKey)

			// 获取当前时间
			milli := time.Now().Format(constant.TimeLayout)

			// 将 被强制挤下线的用户，以及时间，保存到缓存中，提示给前端用户！
			SetString(constant.ForcedOfflineKey+oldLoginId, milli, time.Duration(conf.PromptExpired)*time.Second)
		}
		// 设置 在线用户，为新用户
		SetString(onlineUserKey, loginId, time.Duration(conf.TokenExpired)*time.Second)
	}
}

// SetRefreshToken 设置 RefreshToken
func SetRefreshToken(loginId string, token *domain.TokenValue) {
	jsonStr, err := json.Marshal(token)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	SetString(constant.RefreshTokenKey+loginId, string(jsonStr), time.Duration(conf.TokenExpired)*time.Second)
}

// GetAccessToken 获取 AccessToken
func GetAccessToken(loginId string) (*domain.TokenValue, *exception.OAuth2Error) {
	return getToken(constant.LoginTokenKey + loginId)
}

// GetRefreshToken 获取 RefreshToken
func GetRefreshToken(loginId string) (*domain.TokenValue, *exception.OAuth2Error) {
	return getToken(constant.RefreshTokenKey + loginId)
}

func getToken(loginId string) (*domain.TokenValue, *exception.OAuth2Error) {
	str := GetString(loginId)
	if len(str) < 1 {
		return nil, exception.AccessTokenInvalid
	}
	log.Logger.Debug("getToken JSON Str : ", str)
	var token domain.TokenValue
	err := json.Unmarshal([]byte(str), &token)
	if err != nil {
		return nil, &exception.OAuth2Error{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return &token, nil
}

// SetCodeModeParam 设置授权码模式中，请求参数
func SetCodeModeParam(codeId string, param *params.GetCodeModeParam) {
	jsonStr, err := json.Marshal(param)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	SetString(constant.CodeModeKey+codeId, string(jsonStr), time.Duration(conf.CodeExpired)*time.Second)
}

// GetCodeModeParam 获取 授权码模式中，请求参数
func GetCodeModeParam(codeId string) (*params.GetCodeModeParam, *exception.OAuth2Error) {
	str := GetString(constant.CodeModeKey + codeId)
	if len(str) < 1 {
		return nil, exception.CodeInvalid
	}
	var param params.GetCodeModeParam
	err := json.Unmarshal([]byte(str), &param)
	if err != nil {
		return nil, &exception.OAuth2Error{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return &param, nil
}

// SetLoginUserInfo 设置 登录成功的用户信息
func SetLoginUserInfo(code string, user *domain.UserInfo) {
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	SetString(constant.CodeLoginUserKey+code, string(jsonStr), time.Duration(conf.CodeExpired)*time.Second)
}

// GetLoginUserInfo 获取 登录成功的用户信息
func GetLoginUserInfo(code string) (*domain.UserInfo, *exception.OAuth2Error) {
	str := GetString(constant.CodeLoginUserKey + code)
	if len(str) < 1 {
		return nil, exception.CodeInvalid
	}
	var user domain.UserInfo
	err := json.Unmarshal([]byte(str), &user)
	if err != nil {
		return nil, &exception.OAuth2Error{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return &user, nil
}

// 获取 string 的 key

func GetString(key string) (str string) {
	switch cacheType {
	case "redis":
		data, err := redisCache.Get(ctx, key).Result()
		if err != nil || len(data) < 1 {
			log.Logger.Error(err)
		}
		str = data
		break
	default:
		if data, ok := goCache.Get(key); ok {
			str = util.Strval(data)
		}
		break
	}
	return
}

// 设置 string 的 key

func SetString(key, val string, expired time.Duration) {
	switch cacheType {
	case "redis":
		_, err := redisCache.Set(ctx, key, val, expired).Result()
		if err != nil {
			log.Logger.Error(err)
		}
		break
	default:
		goCache.Set(key, val, expired)
		break
	}
	return
}

// Remove 移除 key
func Remove(key string) {
	switch cacheType {
	case "redis":
		_, err := redisCache.Del(ctx, key).Result()
		if err != nil {
			log.Logger.Error(err)
		}
		break
	default:
		goCache.Delete(key)
		break
	}
	return
}
