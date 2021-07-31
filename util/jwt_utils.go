package util

import (
	"encoding/json"
	"gin_oauth2_server/config"
	"gin_oauth2_server/constant"
	"gin_oauth2_server/exception"
	"gin_oauth2_server/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

var conf = config.Config.OAuth2Server

var jwtSecret = []byte(conf.Secret)

// CreateAccessToken  创建 access_token
func CreateAccessToken(loginId string) (string, error) {
	return createToken(loginId, conf.TokenExpired)
}

// CreateRefreshToken 创建 刷新 token
func CreateRefreshToken(loginId string) (string, error) {
	return createToken(loginId, conf.RefreshTokenExpired)
}

// CreateLoginId 创建 随机的 loginId
func CreateLoginId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// createToken 创建 token
func createToken(loginId string, expired int64) (string, error) {
	expiredInt := time.Now().Add(time.Duration(expired) * time.Second)
	// 过期时间秒
	expiredExpired := strconv.FormatInt(expiredInt.Unix(), 10)
	// 过期时间秒
	claims := jwt.MapClaims{
		"loginId": loginId,
		"expired": expiredExpired,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	log.Logger.Debug("创建 过期时间: ", expiredInt.Format(constant.TimeLayout))
	return token, err
}

// GetIdFromToken 获取 token 中的 tokenId
func GetIdFromToken(bearerToken string) (string, *exception.OAuth2Error) {
	// token 前缀必须包含 "Bearer "
	if !strings.HasPrefix(bearerToken, conf.TokenType) {
		return "", exception.TokenFormatError
	}
	jwtToken := strings.Trim(bearerToken[len(conf.TokenType):], " ")
	tokenClaims, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", &exception.OAuth2Error{Status: 500, Message: err.Error()}
	}
	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if ok && tokenClaims.Valid {
		loginId := Strval(claims["loginId"])
		expired, err := strconv.ParseInt(Strval(claims["expired"]), 10, 64)
		// 如果当前时间大于过期时间，那么 token 就是过期了
		if err != nil || time.Now().Unix() >= expired {
			return "", exception.AccessTokenExpired
		}
		//time.Unix  是time包里的函数，将时间戳转为Time类型
		log.Logger.Debug("比较 过期时间: ", time.Unix(expired, 0).Format(constant.TimeLayout))

		return loginId, nil
	}
	return "", exception.TokenFormatError
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
