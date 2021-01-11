package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/model"
	"time"

	"gorm.io/gorm"
)

// SetUpDB method
func SetUpDB() {

	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))

	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	migration(db)
}

func migration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
