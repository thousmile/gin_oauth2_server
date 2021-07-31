package constant

const (
	CaptchaCodeKey = "captcha_codes:" // 验证码 key

	LoginTokenKey = "login_tokens:" // 登录 token key

	RefreshTokenKey = "refresh_tokens:" //refresh tokens key

	OnlineUserKey = "online_user:" // 在线用户，令牌前缀

	ForcedOfflineKey = "forced_offline_user:" // 强制下线，令牌前缀

	SmsKey = "sms_codes:" // 短信验证码， 令牌前缀

	TencentKey = "tencent_state:" // 腾讯qq和微信 state字段 令牌前缀

	ClientKey = "clients:" // 客户端

	CodeModeKey = "code_mode:" // 授权码模式 code 参数

	CodeLoginUserKey = "code:" // 授权码模式 返回给第三方应用的 code 关联登录用户的信息

	CodeClientIdKey = "code_clientId:" // 授权码模式 当用户登录成功后

	TimeLayout = "2006-01-02 15:04:05" // 时间格式化
)

const (
	Code              = "code"               // 授权码模式，1.获取code
	AuthorizationCode = "authorization_code" // 授权码模式，2.通过code获取access_token
	Password          = "password"           // 密码模式
	ClientCredentials = "client_credentials" // 客户端模式
	WeChat            = "we_chat"            // 微信登陆模式
	TencentQq         = "tencent_qq"         // 腾讯 QQ 登陆模式
	Sms               = "sms"                // 手机短信模式
)

const (
	Disable = iota // 禁用
	Normal         // 正常
	Locking        // 锁定
)

const (
	Female  = iota // 女
	Male           // 男
	Unknown        // 未知
)

const (
	NO  = iota // 是
	YES        // 否
)
