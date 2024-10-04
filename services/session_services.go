package services

import (
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/repository"
	"time"
)

func BookGameSession(sessionRequest *models.SessionRequest)error{
	return repository.BookGameSession(sessionRequest)
}

func CheckGameAvailability(gameID int,startTime,endTime time.Time)(bool,error){
	return repository.CheckGameAvailability(gameID,startTime,endTime)
}

func CancelSession(sessionID string)error{
	return repository.CancelSession(sessionID)
}

func RescheduleSession(sessionID string, rescheduleRequest *models.SessionRescheduleRequest) error {
	return repository.RescheduleSession(sessionID, rescheduleRequest)
}
