package handlers

import (
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/services"
	"mows-game-center-time-mgt/utils/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func BookGameSession(c *gin.Context) {
	var sessionRequest models.SessionRequest
	if err := c.ShouldBindJSON(&sessionRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid session data", nil, err.Error()))
		return
	}
	//Check if double booking
	gameIDint, _ := strconv.Atoi(sessionRequest.GameID)
	available, err := services.CheckGameAvailability(gameIDint, sessionRequest.StartTime, sessionRequest.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to check availability", nil, err.Error()))
		return
	}
	if !available {
		c.JSON(http.StatusConflict, response.ClientResponse(http.StatusConflict, "Failed to book session", nil, nil))
		return
	}
	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "Game session booked successfully", sessionRequest, nil))
}

// GameAvailability
func CheckAvailableGames(c *gin.Context) {
	gameID := c.Param("game_id")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid start time format", nil, err.Error()))
		return
	}
	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid end time format", nil, err.Error()))
		return
	}
	gameIDint, _ := strconv.Atoi(gameID)
	avilable, err := services.CheckGameAvailability(gameIDint, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to check availability", nil, err.Error()))
		return
	}

	if avilable {
		c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game slot is available", nil, nil))
	} else {
		c.JSON(http.StatusConflict, response.ClientResponse(http.StatusConflict, "Game slot is booked", nil, nil))
	}
}

// Cancel the session
func CancelSession(c *gin.Context) {
	sessionID := c.Param("id")

	if err := services.CancelSession(sessionID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to cancel session", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Session canceled successfully", nil, nil))
}

// Reschedule the session
func RescheduleSession(c *gin.Context) {
	sessionID := c.Param("id")
	var rescheduleRequest models.SessionRescheduleRequest
	if err := c.ShouldBindJSON(&rescheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid reshcedule data", nil, err.Error()))
		return
	}
	//find available new slot
	available, err := services.CheckGameAvailability(rescheduleRequest.GameID, rescheduleRequest.NewStartTime, rescheduleRequest.NewEndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to check availability", nil, err.Error()))
		return
	}

	if !available {
		c.JSON(http.StatusConflict, response.ClientResponse(http.StatusConflict, "Game slot already booked", nil, "Double booking detected"))
	}

	if err = services.RescheduleSession(sessionID, &rescheduleRequest); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to reschedule session", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Session rescheduled successfully", rescheduleRequest, nil))
}

// user starts a game and time is tracked
func StartGameSession(c *gin.Context) {
	var sessionRequest models.SessionRequest
	if err := c.ShouldBindJSON(&sessionRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid session data", nil, err.Error()))
		return
	}

	//if the user is already in active session for the same game
	activeSession, err := services.CheckActiveGameSession(sessionRequest.GameID, sessionRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to check active session", nil, err.Error()))
		return
	}
	if activeSession != nil {
		c.JSON(http.StatusConflict, response.ClientResponse(http.StatusConflict, "User already has an active session for this game", nil, "Active session detected"))
		return
	}

	sessionRequest.StartTime = time.Now()

	err := services.StartGameSession(&sessionRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Failed to start game session", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Game session started successfully", sessionRequest, nil))
}
