package repository

import (
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
	"time"

	"gorm.io/gorm"
)

// Get All user details
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	//Get all users
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Delete User By ID from the DB
func DeleteUser(id int) error {
	if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func CheckUserAvailabilityWithID(userID int) bool {
	var count int64
	err := db.DB.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

// Retrieves a user by their ID
func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Add new user
func AddNewUser(user *models.User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(id int, updateUser *models.User) error {
	if err := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(updateUser).Error; err != nil {
		return err
	}
	return nil
}

// (Admin forcibly ends a game session)
func AdminEndGameSession(db *gorm.DB, sessionID string, endTime time.Time) error {
	//Find the session by ID
	var session models.Session
	if err := db.First(&session, "id = ?", sessionID).Error; err != nil {
		return err
	}
	session.EndTime = endTime
	session.Status = "completed"
	return db.Save(&session).Error
}
