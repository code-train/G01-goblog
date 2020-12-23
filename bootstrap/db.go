package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

// SetUpDB method
func SetUpDB() {

	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
