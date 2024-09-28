package services

import (
	"errors"
	"mows-game-center-time-mgt/utils/models"
)

var sessions []models.Session

// Start a new game session
func StartSession(newSession models.Session) models.Session {
	newSession.ID = len(sessions) + 1
	newSession.Status = "ongoing"
	sessions = append(sessions, newSession)
	return newSession
}

// End a session
func EndSession(id int) (models.Session, error) {
	for i, session := range sessions {
		if session.ID == id && session.Start == "ongoing" {
			sessions[i].Status = "completed"
			//Calculate the total price (game duration * rate)
			sessions[i].TotalPrice = calculateSessionPrice(session.GameID, session.Start, session.End)
			return sessions[i], nil
		}
	}
	return models.Session{}, errors.New("session not found or already ended")
}

func calculateSessionPrice(gameID int, start, end string) float64 {

}


func GetSessionsForUser(userID int) []models.Session{
	var userSessions []models.Session
	for _,session := range sessions{
		if session.UserID == userID{
			userSessions = append(userSessions, session)
		}
	}
	return userSessions
}

