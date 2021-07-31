package mapper

import (
	"fmt"
	conf "gin_oauth2_server/config"
	"gin_oauth2_server/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

var BaseMapper *gorm.DB

func init() {
	dbType := strings.ToLower(conf.Config.DbType)
	var err error
	switch dbType {
	case "mysql":
		mc := conf.Config.Mysql
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mc.Username, mc.Password, mc.Host, mc.Port, mc.DbName,
		)
		log.Logger.Debug("mysql : ", dsn)
		BaseMapper, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql":
		pc := conf.Config.Postgre
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			pc.Host, pc.Username, pc.Password, pc.DbName, pc.Port,
		)
		log.Logger.Debug("PostgreSQL : ", dsn)
		BaseMapper, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver":
		ms := conf.Config.SqlServer
		// github.com/denisenkom/go-mssqldb
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			ms.Username, ms.Password, ms.Host, ms.Port, ms.DbName,
		)
		log.Logger.Debug("SqlServer : ", dsn)
		BaseMapper, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	default:
		log.Logger.Debug("SQLite : ", conf.Config.SQLite.DbPath)
		BaseMapper, err = gorm.Open(sqlite.Open(conf.Config.SQLite.DbPath), &gorm.Config{})
	}
	if err != nil {
		print(err)
		os.Exit(0)
	}
	dbPool, err := BaseMapper.DB()
	if err != nil {
		print(err)
		os.Exit(0)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbPool.SetMaxIdleConns(3)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbPool.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbPool.SetConnMaxLifetime(time.Hour)
	// 连接池里面的连接最大空闲时长。
	dbPool.SetConnMaxIdleTime(time.Hour)

}
