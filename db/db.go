package db

import (
	"fmt"
	"log"
	"mows-game-center-time-mgt/config"
	"mows-game-center-time-mgt/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the connection to the database and runs migrations
func ConnectDatabase(confg config.Config) (*gorm.DB, error) {
	connectTo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", confg.DBHost, confg.DBUser, confg.DBName, confg.DBPort, confg.DBPassword)
	//Open the connection
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database:%w", err)
	}
	//Assign the DB variable for global use
	DB = db
	runMigrations(db)

	return DB, nil
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Game{},
		&models.User{},
		&models.Session{},
		//&models.Admin{}, // Ensure this model is properly defined
		//&models.Payment{}, // Uncomment if defined correctly
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	} else {
		log.Println("Database migrations completed successfully.")
	}
}
