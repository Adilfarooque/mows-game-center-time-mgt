package services

import (
	"errors"
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/repository"
	"time"
)

func BookGameSession(sessionRequest *models.SessionRequest) error {
	return repository.BookGameSession(sessionRequest)
}

func CheckGameAvailability(gameID int, startTime, endTime time.Time) (bool, error) {
	return repository.CheckGameAvailability(gameID, startTime, endTime)
}

func CancelSession(sessionID string) error {
	return repository.CancelSession(sessionID)
}

func RescheduleSession(sessionID string, rescheduleRequest *models.SessionRescheduleRequest) error {
	return repository.RescheduleSession(sessionID, rescheduleRequest)
}

func StartGameSession(sessionRequest *models.SessionRequest) error {
	return repository.StartGameSession(sessionRequest)
}

func CheckActiveGameSession(gameID string, UserID string) (*models.Session, error) {
	return repository.GetActiveGameSession(gameID, UserID)
}

func GetGameSessionByID(sessionID string) (*models.Session, error) {
	var session models.Session
	result := db.DB.First(&session, sessionID)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, errors.New("session not found")
		}
		return nil, result.Error
	}
	return &session, nil
}

func EndGameSession(sessionID string,endTime time.Time,duration time.Duration)error{
	return repository.EndGameSession(sessionID,endTime,duration)
}