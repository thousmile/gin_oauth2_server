package domain

// OAuth2Token OAuth2.0 token 返回
type OAuth2Token struct {
	/**
	 * 表示访问令牌，必选项。
	 * @date 2020/7/25 15:02
	 */
	AccessToken string `json:"access_token"`

	/**
	 * 表示令牌类型，该值大小写不敏感，必选项，可以是bearer类型或mac类型。
	 * @date 2020/7/25 15:02
	 */
	TokenType string `json:"token_type"`

	/**
	 * 表示更新令牌，用来获取下一次的访问令牌，可选项。
	 * @date 2020/7/25 15:02
	 */
	RefreshToken string `json:"refresh_token"`

	/**
	 * 表示权限范围，如果与客户端申请的范围一致，此项可省略。
	 * @date 2020/7/25 15:02
	 */
	Scope string `json:"scope"`

	/**
	 * 表示过期时间，单位为秒。如果省略该参数，必须其他方式设置过期时间。
	 * @date 2020/7/25 15:02
	 */
	ExpiresIn int64 `json:"expires_in"`
}
