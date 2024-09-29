package services

import (
	"errors"
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/repository"
)

var games []models.Games

func GetAllGames() ([]models.Games, error) {
	return repository.GetAllGames()
}
func AddGame(newGame models.Games) models.Games {
	newGame.ID = len(games) + 1
	games = append(games, newGame)
	return newGame
}

func UpdateGame(id int, updateGame models.Games) (models.Games, error) {
	var game models.Games
	//Find game by ID
	if err := db.DB.First(&game, id).Error; err != nil {
		return models.Games{}, errors.New("game not found")
	}

	game.Name = updateGame.Name
	game.Category = updateGame.Category
	game.Rating = updateGame.Rating

	if err := db.DB.Save(&game).Error; err != nil {
		return models.Games{}, err
	}
	return game, nil
}

// Remove Game by ID
func RemoveGame(id int) error {
	for i, game := range games {
		if game.ID == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return errors.New("game not found")
}

func GetgameByID(id int) (models.Games, error) {
	for _, game := range games {
		if game.ID == id {
			return game, nil
		}
	}
	return models.Games{}, errors.New("game not found")
}

func GetGamesByCategory(category string) []models.Games {
	var filteredGames []models.Games
	for _, game := range games {
		if game.Category == category {
			filteredGames = append(filteredGames, game)
		}
	}
	return filteredGames
}

func RateGame(id int, rating float64) (models.Games, error) {
	for i, game := range games {
		if game.ID == id {
			games[i].Rating = rating
			return games[i], nil
		}
	}
	return models.Games{}, errors.New("game not found")
}
