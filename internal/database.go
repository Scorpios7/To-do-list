package database

import (
	"fmt"

	"github.com/Scorpios7/To-do-list/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() error {
	dsn := viper.GetString("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("fail to connect database: %v", err)
	}
	err = database.AutoMigrate(&models.Todo{})
	if err != nil {
		return fmt.Errorf("fail to migrate todo table: %v", err)
	}
	db = database
	fmt.Println("connect and migrate database successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}
