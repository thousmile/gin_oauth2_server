package config

import (
	"gin_oauth2_server/log"
	"github.com/spf13/viper"
	"os"
)

// Config 全局配置配置文件
var Config *AppConfig

func init() {
	viper.SetDefault("AppName", "GIN OAUTH2.0 SERVER")
	viper.SetDefault("Mode", "debug")
	viper.SetDefault("Profile", "dev")
	viper.SetDefault("Port", 9018)
	viper.SetDefault("DbType", "SQLite")
	viper.SetDefault("CacheType", "Map")

	// SQLite 默认配置
	viper.SetDefault("SQLite.DbPath", "gin_oauth2_server.db3")

	// redis 默认配置
	viper.SetDefault("Redis", &RedisConfig{
		Host: "localhost",
		Port: 6379,
		Db:   0,
	})

	// Mysql 默认配置
	viper.SetDefault("Mysql", &MysqlConfig{
		Host:     "localhost",
		Port:     3306,
		DbName:   "gin_oauth2_server",
		Username: "root",
		Password: "root",
	})

	// PostgreSQL 默认配置
	viper.SetDefault("Postgre", &PostgreSQLConfig{
		Host:     "localhost",
		Port:     5432,
		DbName:   "gin_oauth2_server",
		Username: "postgre",
		Password: "postgre",
	})

	// SqlServer 默认配置
	viper.SetDefault("SqlServer", &SqlServerConfig{
		Host:     "localhost",
		Port:     1433,
		DbName:   "gin_oauth2_server",
		Username: "sa",
		Password: "123456",
	})

	viper.SetDefault("OAuth2Server", &OAuth2ServerConfig{
		TokenExpired:        7200,
		RefreshTokenExpired: 14400,
		SmsCodeExpired:      600,
		CodeExpired:         600,
		PromptExpired:       600,
		TokenType:           "Bearer ",
		TokenHeader:         "Authorization",
		Secret:              "2N321lIkh$*5!YZykD$7@ApaM8rNt4&5!YZyk8rNt4&5!YZykD$7@ApaM@!IfNt4&5!YZykD$7@ApD$7@&4CZ7eqKe!s",
		Sso:                 true,
	})
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")         // 查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname")        // 多次调用AddConfigPath，可以添加多个搜索路径
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	viper.AddConfigPath("../")                   // optionally look for config in the working directory
	viper.AddConfigPath("./conf/")               // 还可以在工作目录中搜索配置文件
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Logger.Panicf("Fatal error config file: %v \n", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Logger.Panicf("Fatal error config file: %v \n", err)
		os.Exit(1)
	}
}

type AppConfig struct {
	AppName string

	Version string

	Profile string

	Port int

	Mode string

	DbType string

	CacheType string

	SQLite SQLiteConfig

	Redis RedisConfig

	Mysql MysqlConfig

	Postgre PostgreSQLConfig

	SqlServer SqlServerConfig

	OAuth2Server OAuth2ServerConfig
}

// SQLiteConfig SQLite配置
type SQLiteConfig struct {
	DbPath string
}

// RedisConfig redis 配置
type RedisConfig struct {
	Host string

	Port int

	Db int

	Password string
}

// MysqlConfig mysql 配置
type MysqlConfig struct {
	Host string

	Port int

	DbName string

	Username string

	Password string
}

// PostgreSQLConfig Postgre 配置
type PostgreSQLConfig struct {
	Host string

	Port int

	DbName string

	Username string

	Password string
}

// SqlServerConfig SqlServer 配置
type SqlServerConfig struct {
	Host string

	Port int

	DbName string

	Username string

	Password string
}

// OAuth2ServerConfig  OAuth2.0 Server 配置
type OAuth2ServerConfig struct {

	/**
	 * token 过期时间，单位(秒)，默认 2 小时
	 */
	TokenExpired int64

	/**
	 * refreshToken 过期时间，单位(秒)， 默认 3 小时
	 */
	RefreshTokenExpired int64

	/**
	 * 短信验证码过期时间 单位(秒)
	 */
	SmsCodeExpired int64

	/**
	 * 授权码的 code 过期时间 单位(秒)
	 */
	CodeExpired int64

	/**
	 * 用户被挤下线，提示的过期时间 单位(秒)
	 */
	PromptExpired int64

	/**
	 * token 类型
	 */
	TokenType string

	/**
	 * token 在请求头中的名称
	 */
	TokenHeader string

	/**
	 * 秘钥
	 */
	Secret string

	/**
	 * 单点登录，是否启用
	 */
	Sso bool
}
