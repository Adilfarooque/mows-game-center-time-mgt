package repository

import (
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
)

var games []models.Games

func GetAllGames() ([]models.Games, error) {
	err := db.DB.Raw("SELECT * FROM games").Scan(&games).Error
	if err != nil {
		return nil, err
	}
	return games, nil
}
