package db

import (
	"fmt"
	"mows-game-center-time-mgt/config"
	"mows-game-center-time-mgt/utils/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(confg config.Config) (*gorm.DB, error) {
	connectTo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", confg.DBHost, confg.DBUser, confg.DBName, confg.DBPort, confg.DBPassword)
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database:%w", err)
	}
	DB = db
	DB.AutoMigrate(models.Games{})
	DB.AutoMigrate(models.Payment{})
	DB.AutoMigrate(models.Session{})
	DB.AutoMigrate(models.User{})
	return DB, err
}
