package main

import (
	"fmt"
	conf "gin_oauth2_server/config"
	handlers "gin_oauth2_server/handler"
	"gin_oauth2_server/log"
	"gin_oauth2_server/middleware"
	"github.com/gin-gonic/gin"
)

func GinServerStart() {
	gin.SetMode(conf.Config.Mode)
	r := gin.New()

	// 注册zap logger 相关中间件
	r.Use(log.GinLogger(), log.GinRecovery(true))

	// 如需跨域可以打开
	r.Use(middleware.Cors())

	// 引入 静态资源
	r.Static("/assets", "./static")
	// 设置 favicon.ico
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 引入 模板引擎
	r.LoadHTMLGlob("templates/*")

	// 首页
	r.GET("/", handlers.IndexHandler)

	// 授权码模式，第一步，第三方应用，请求获取用户信息。在此校验客户端是否正确
	r.GET("/code", handlers.GetCodeHandler)

	// 授权码模式，第二步，用户登录
	r.POST("/login", handlers.LoginHandler)

	// 授权码模式，第三步，第三方应用通过 code 来换取 access_token
	r.POST("/access_token", handlers.GetAccessToken)

	// 初始化几个客户端
	r.GET("/init/clients", handlers.ClientInitDataHandler)

	r.GET("/init/users", handlers.UserInitHandler)

	// 密码模式
	r.POST("/password", handlers.PasswordHandler)

	// 客户端模式
	r.POST("/client", handlers.ClientHandler)

	// 发送 短信验证码
	r.POST("/sms/send", handlers.SendSmsHandler)

	// 短信验证码 模式
	r.POST("/sms", handlers.SmsHandler)

	// 刷新 token
	r.POST("/refresh", handlers.RefreshHandler)

	// 授权认证
	r.Use(middleware.Authorize())

	// 退出登录
	r.POST("/logout", handlers.LogoutHandler)

	// 登录的用户信息
	r.GET("/loginInfo", handlers.LoginInfoHandler)

	// listen and serve on 0.0.0.0:9018
	addr := fmt.Sprintf(":%v", conf.Config.Port)
	r.Run(addr)
}
