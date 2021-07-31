package exception

var (
	/**
	 * 认证错误，不知道啥原因
	 */
	OAuth2Exception = &OAuth2Error{400004, "认证错误!"}

	/**
	 * access_token 不存在
	 */
	AccessTokenInvalid = &OAuth2Error{400010, "access_token 不存在"}

	/**
	 * access_token 过期
	 */
	AccessTokenExpired = &OAuth2Error{400011, "access_token 已过期"}

	/**
	 * token 格式错误
	 */
	TokenFormatError = &OAuth2Error{400012, "access_token 格式错误"}

	// refresh_token 不存在
	RefreshTokenInvalid = &OAuth2Error{400014, "refresh_token 不存在"}

	// RefreshTokenExpired refresh_token 过期
	RefreshTokenExpired = &OAuth2Error{400015, "refresh_token 已过期"}

	// TokenUserInvalid 此服务，必须是非客户端模式，授权才可以调用
	TokenUserInvalid = &OAuth2Error{400016, "此服务，必须是非客户端模式，授权才可以调用"}

	// UserInvalid 用户不存在
	UserInvalid = &OAuth2Error{400020, "用户不存在"}

	// UserPasswordError 用户密码错误
	UserPasswordError = &OAuth2Error{400021, "用户密码错误"}

	/**
	 * 当前用户被禁用
	 */
	UserDisable = &OAuth2Error{400022, "当前用户被禁用"}

	/**
	 * 当前用户被锁定
	 */
	UserLocking = &OAuth2Error{400023, "当前用户被禁用"}

	/**
	 * 用户手机号不存在
	 */
	UserMobileInvalid = &OAuth2Error{400024, "用户手机号不存在"}

	/**
	 * 客户端不存在
	 */
	ClientInvalid = &OAuth2Error{400030, "客户端不存在"}

	/**
	 * 客户端被禁用
	 */
	ClientDisable = &OAuth2Error{400031, "客户端被禁用"}

	/**
	 * 客户端秘钥错误
	 */
	ClientSecretError = &OAuth2Error{400032, "客户端秘钥错误"}

	/**
	 * 授权类型错误
	 */
	AuthorizationGrantType = &OAuth2Error{400033, "授权类型错误"}

	/**
	 * 授权码模式中，使用code来换取 access_token，这个code不存在。
	 */
	CodeInvalid = &OAuth2Error{400034, "code不存在"}

	/**
	 * 授权码模式中，回调的域名和数据库的域名不一致
	 */
	DomainNameIllegal = &OAuth2Error{400035, "回调的域名与设置的域名不一致"}

	/**
	 * 客户端不相同
	 */
	ClientIsDifferent = &OAuth2Error{400036, "两次认证，客户端不一致"}

	/**
	 * 请求参数解析错误
	 */
	RequestParamValidate = &OAuth2Error{400044, "请求参数解析错误"}

	/**
	 * 验证码错误
	 */
	VerificationCodeError = &OAuth2Error{400045, "验证码错误"}

	/**
	 * 短信 验证码错误
	 */
	SmsCodeError = &OAuth2Error{400046, "短信验证码错误"}
)

type OAuth2Error struct {
	Status int

	Message string
}
