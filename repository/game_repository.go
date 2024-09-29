package repository

import (
	"errors"
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
)

var games []models.Game

func AddNewGame(newGame models.Game) (models.Game, error) {
	//Insert the new game into db
	if err := db.DB.Create(&newGame).Error; err != nil {
		return models.Game{}, err
	}
	return newGame, nil
}

func GetAllGames() ([]models.Game, error) {
	if err := db.DB.Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func GetGameByID(id int) (models.Game, error) {
	var game models.Game
	if err := db.DB.First(&game, id).Error; err != nil {
		return game, err
	}
	return game, nil
}

func UpdateGame(updateGame models.Game) error {
	result := db.DB.Save(&updateGame)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func RemoveGame(id int) error {
	var game models.Game

	if err := db.DB.First(&game, id).Error; err != nil {
		return errors.New("game not found")
	}
	if err := db.DB.Delete(&game).Error; err != nil {
		return err
	}

	return nil
}
