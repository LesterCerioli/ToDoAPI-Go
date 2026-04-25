package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("📦 Connected to SQLite database (pure-Go)")
}

func AutoMigrate(models ...interface{}) {
	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("✅ Database migrated")
}
