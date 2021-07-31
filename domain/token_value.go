package domain

import (
	"encoding/json"
	"gin_oauth2_server/log"
	"strings"
)

type TokenValue struct {
	/**
	 * 全局唯一认证 Id
	 */
	LoginId string `json:"loginId"`

	/**
	 * 认证授权方式
	 */
	GrantType string `json:"grantType"`

	/**
	 * 客户端详情
	 */
	Client *ClientDetails `json:"client"`

	/**
	 * 用户信息
	 */
	User *UserInfo `json:"user"`

	/**
	 * 登录时间
	 */
	LoginTime int64 `json:"loginTime"`
}

type ClientDetails struct {

	/**
	 * 客户端 唯一认证 Id
	 */
	ClientId string `gorm:"primaryKey" json:"clientId"`

	/**
	 * 客户端 密钥
	 */
	Secret string `json:"-"`

	/**
	 * '客户端名称'
	 */
	Name string `json:"name"`

	/**
	 * 客户端 图标
	 */
	Logo string `json:"logo"`

	/**
	 * 客户端 描述
	 */
	Description string `json:"description"`

	/**
	 * 客户端类型
	 */
	ClientType uint8 `json:"clientType"`

	/**
	 * 授权类型 json 数组格式
	 */
	GrantTypes string `json:"grantTypes"`

	/**
	 * 域名地址，如果是 授权码模式，
	 * 必须校验回调地址是否属于此域名之下
	 */
	DomainName string `json:"domainName"`

	/**
	* 授权作用域
	 */
	Scope string `json:"scope"`

	/**
	 * 状态 【0.禁用 1.正常 2.锁定 】
	 */
	Status uint8 `json:"status"`
}

// ContainsGrantType 判断是否 包含 授权类型
func (r *ClientDetails) ContainsGrantType(grantType string) bool {
	lower := strings.ToLower(r.GrantTypes)
	var grantTypes []string
	err := json.Unmarshal([]byte(lower), &grantTypes)
	if err != nil {
		panic(err)
	}
	log.Logger.Debugf("ClientId: %s Have GrantTypes: %s", r.ClientId, grantTypes)
	if grantTypes == nil || len(grantTypes) < 1 {
		return false
	}
	for _, g := range grantTypes {
		if strings.EqualFold("*", g) {
			return true
		}
		if strings.EqualFold(g, grantType) {
			return true
		}
	}
	return false
}

type UserInfo struct {
	/**
	 * 用户ID
	 */
	UserId string `gorm:"primaryKey" json:"userId"`

	/**
	 * 头像
	 */
	Avatar string `json:"avatar"`

	/**
	 * 用户名
	 */
	Username string `json:"username"`

	/**
	 * 手机号
	 */
	Mobile string `json:"mobile"`

	/**
	 * 邮箱
	 */
	Email string `json:"email"`

	/**
	 * 用户名称
	 */
	Nickname string `json:"nickname"`

	/**
	 * 密码
	 */
	Password string `json:"-"`

	/**
	 * 状态 【0.禁用 1.正常 2.锁定 】
	 */
	Status uint8 `json:"status"`

	/**
	 * 当前用户是不是管理员 0.否  1. 是
	 */
	AdminFlag uint8 `json:"adminFlag"`
}
