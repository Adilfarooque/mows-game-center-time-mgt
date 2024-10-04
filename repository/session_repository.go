package repository

import (
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
	"time"
)

func BookGameSession(sessionRequest *models.SessionRequest) error {
	session := models.Session{
		GameID:    sessionRequest.GameID,
		UserID:    sessionRequest.UserID,
		StartTime: sessionRequest.StartTime,
		EndTime:   sessionRequest.EndTime,
	}
	if err := db.DB.Create(&session).Error; err != nil {
		return err
	}
	return nil
}

func CheckGameAvailability(gameID int, startTime, endTime time.Time) (bool, error) {
	var count int64
	db.DB.Model(&models.Session{}).Where("game_id = ? AND start_time < ? end_time > ?", gameID, endTime, startTime).Count(&count)
	return count == 0, nil
}

func CancelSession(sessionID string) error {
	return db.DB.Where("id = ?", sessionID).Delete(&models.Session{}).Error
}

func RescheduleSession(sessionID string, rescheduleRequest *models.SessionRescheduleRequest) error {
	updates := map[string]interface{}{
		"start_time": rescheduleRequest.NewStartTime,
		"end_time":   rescheduleRequest.NewEndTime,
	}
	return db.DB.Model(&models.Session{}).Where("id = ?", sessionID).Updates(updates).Error
}
