package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/model"
	"time"

	"gorm.io/gorm"
)

// SetUpDB method
func SetUpDB() {

	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)

	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migraton(db)
}

func migraton(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
