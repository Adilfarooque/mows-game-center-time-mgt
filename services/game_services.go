package services

import (
	"errors"
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/repository"
)

//var games []models.Game

func AddNewGame(newGame models.Game) (models.Game, error) {
	addedGame, err := repository.AddNewGame(newGame)
	if err != nil {
		return models.Game{}, err
	}
	return addedGame, nil
}

func GetgameByID(id int) (models.Game, error) {
	return repository.GetGameByID(id)
}

func GetAllGames() ([]models.Game, error) {
	return repository.GetAllGames()
}

func UpdateGameByID(id int, updateGame models.Game) error {
	//Retrieve the game from the database
	game, err := repository.GetGameByID(id)
	if err != nil {
		return err
	}
	if game.ID == 0 {
		return errors.New("game not found")
	}
	game.Title = updateGame.Title
	game.Category = updateGame.Category
	game.Rating = updateGame.Rating
	game.Popularity = updateGame.Popularity
	game.Price = updateGame.Price
	game.ImageURL = updateGame.ImageURL
	if err := repository.UpdateGame(game); err != nil {
		return err
	}
	return nil
}

func RemoveGame(id int) error {
	return repository.RemoveGame(id)
}

func GetGameByName(title string) (models.Game, error) {
	var game models.Game
	err := db.DB.Where("title = ?", title).First(&game).Error
	if err != nil {
		return models.Game{}, err
	}
	return game, nil
}
