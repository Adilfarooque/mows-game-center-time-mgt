package handlers

import (
	"mows-game-center-time-mgt/services"
	"mows-game-center-time-mgt/utils/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var games []models.Games

// Get All games
func GetAllGames(c *gin.Context) {
	c.JSON(http.StatusOK, games)
}

// Get games by name
func GetGamesByName(c *gin.Context) {
	name := c.Param("name")
	for _, game := range games {
		if game.Name == name {
			c.JSON(http.StatusOK, game)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Game not found"})
}
// Add new game
func AddNewGame(c *gin.Context) {
	var addGame models.Games
	if err := c.ShouldBindJSON(&addGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid game data"})
		return
	}
	games = services.AddGame(addGame)
	c.JSON(http.StatusCreated, gin.H{"message":"Game added successfully","game":addGame})
}
//Update Existing 
//func UpdateGame(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, game := range games {
		if game.ID == id {
			if err := c.ShouldBindJSON(&games[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, games[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Game not found"})
//}

func UpdateGame(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedGame models.Games
    if err := c.ShouldBindJSON(&updatedGame); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid game data"})
        return
    }
    updated, err := services.UpdateGame(id, updatedGame)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Game not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"game": updated})
}


func RemoveGame(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
		err := services.RemoveGame(id)
		if game.ID == id {
			games = append(games[:i], games[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
			return
		}
	c.JSON(http.StatusNotFound, gin.H{"message": "Game not found"})
}

