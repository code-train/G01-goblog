package model

import (
	"fmt"
	config2 "goblog/pkg/config"
	"goblog/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {
	var err error

	var (
		host	= config2.GetString("database.mysql.host")
		port	= config2.GetString("database.mysql.port")
		database = config2.GetString("database.mysql.database")
		username = config2.GetString("database.mysql.username")
		password = config2.GetString("database.mysql.password")
		charset = config2.GetString("database.mysql.charset")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level gormlogger.LogLevel

	if config2.GetBool("app.debug") {
		level = gormlogger.Warn
	} else {
		level = gormlogger.Error
	}

	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})

	logger.LogError(err)
	return DB
}
