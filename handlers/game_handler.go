package handlers

import (
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/services"
	"mows-game-center-time-mgt/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var games []models.Game

// -----------------------------ADMIN SIDE----------------------------//
// Add new game
func AddNewGame(c *gin.Context) {
	var newGame models.Game
	//Bind JSON to game struct
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid game data", nil, err.Error()))
		return
	}
	//add service layer to add the game
	addedGame, err := services.AddNewGame(newGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to add game", nil, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "Game added successfully", addedGame, nil))
}

// Get All games
func GetAllGames(c *gin.Context) {
	//Fetch all games without pagination
	games, err := services.GetAllGames()
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Could't retrieve games", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	//Send Successfull response
	success := response.ClientResponse(http.StatusOK, "Successfully retrieved all games", games, nil)
	c.JSON(http.StatusOK, success)
}

func GetGamesByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid game ID", nil, err.Error()))
		return
	}
	game, err := services.GetgameByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "Game not found", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game retrieved successfully", game, nil))
}

func UpdateGameByID(c *gin.Context) {
	// Get the game ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid game ID", nil, err.Error()))
		return
	}
	// Bind the incoming JSON to the updatedGame struct
	var updatedGame models.Game
	if err := c.ShouldBindJSON(&updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid input", nil, err.Error()))
		return
	}
	// Call the service to update the game by ID
	err = services.UpdateGameByID(id, updatedGame)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "Game not found", nil, err.Error()))
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game updated successfully", updatedGame, nil))
}

func GetGameByName(c *gin.Context) {
	title := c.Param("title")
	game, err := services.GetGameByName(title)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "Game not found", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game retrived successfully", game, nil))
}

func RemoveGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid game ID", nil, err.Error()))
		return
	}
	if err := services.RemoveGame(id); err != nil {
		if err.Error() == "game not found" {
			c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, "Game not found", nil, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to delete game", nil, err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game deleted successfully", nil, nil))
}
