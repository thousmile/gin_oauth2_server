package params

import "regexp"

import (
	"errors"
)

type DefaultModeParam struct {

	/**
	 * 授权类型
	 */
	GrantType string `json:"grant_type" form:"grant_type"`

	/**
	 * 客户端 ID
	 */
	ClientId string `json:"client_id" form:"client_id"`

	/**
	 * 客户端密钥
	 */
	ClientSecret string `json:"client_secret" form:"client_secret"`
}

var pattern = regexp.MustCompile(`(authorization_code|password|client_credentials|tencent_qq|we_chat|sms)`)

func (p *DefaultModeParam) Validator() error {
	if len(p.ClientId) < 10 {
		return errors.New("client_id 格式错误！")
	}
	if len(p.ClientSecret) < 10 {
		return errors.New("client_secret 格式错误！")
	}
	if len(p.ClientSecret) < 4 {
		return errors.New("grant_type 格式错误！")
	}
	if ok := pattern.MatchString(p.GrantType); !ok {
		return errors.New("授权类型，必须是[ authorization_code | password | client_credentials | tencent_qq | we_chat | sms ]之一！")
	}
	return nil
}

// ClientModeParam 客户端模式
type ClientModeParam struct {
	DefaultModeParam
	Scope string `json:"scope" form:"scope"`
}

// PasswordModeParam 密码模式
type PasswordModeParam struct {
	DefaultModeParam
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// SendSmsParam 发送短信验证码
type SendSmsParam struct {

	/**
	 * 登陆的用户手机号
	 */
	Mobile string `json:"mobile" form:"mobile"`

	/**
	 * 客户端 ID
	 */
	ClientId string `json:"client_id" form:"client_id"`
}

// SmsModeParam 短信验证码模式
type SmsModeParam struct {
	DefaultModeParam

	/**
	 * 登陆的用户手机号
	 */
	Mobile string `json:"mobile" form:"mobile"`

	/**
	 * 短信验证码
	 */
	Code string `json:"code" form:"code"`
}

// GetCodeModeParam 授权码模式获取 code 参数
type GetCodeModeParam struct {

	// 表示授权类型，必选项，此处的值固定为"code"
	ResponseType string `json:"response_type" form:"response_type"`

	// 表示客户端的ID，必选项
	ClientId string `json:"client_id" form:"client_id"`

	// 表示 重定向 URI
	RedirectUri string `json:"redirect_uri" form:"redirect_uri"`

	/**
	 * 表示申请的权限范围，可选项，
	 * user_base : 只可以获取用户的 基本信息
	 * user_info  ：获取用户的全部信息
	 */
	Scope string `json:"scope" form:"scope"`

	// 表示客户端的当前状态，可以指定任意值，认证服务器会原封不动地返回这个值。
	State string `json:"state" form:"state"`
}

// CodeModeParam 授权码模式。通过 code 获取 access_token
type CodeModeParam struct {
	DefaultModeParam

	/**
	 * 授权码 的值
	 */
	Code string `json:"code" form:"code"`
}

//UserLoginParam 授权码模式。用户登录
type UserLoginParam struct {

	/**
	 * 一个随机的 授权ID，由系统随机生成，绑定给每个前来授权的第三方应用
	 */
	CodeId string `json:"codeId" form:"codeId"`

	/**
	 * 用户名
	 */
	Username string `json:"username" form:"username"`

	/**
	 * 密码
	 */
	Password string `json:"password" form:"password"`
}
