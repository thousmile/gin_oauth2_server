module gin_oauth2_server

go 1.16

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-redis/redis/v8 v8.11.1
	github.com/google/uuid v1.3.0
	github.com/jameskeane/bcrypt v0.0.0-20120420032655-c3cd44c1e20f
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/spf13/viper v1.8.1
	go.uber.org/zap v1.17.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.7
	gorm.io/gorm v1.21.12
)
