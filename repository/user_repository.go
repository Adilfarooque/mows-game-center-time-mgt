package repository

import (
	"mows-game-center-time-mgt/db"
	"mows-game-center-time-mgt/models"
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

func CheckUserAvailabilityWithID(userID int)bool{
	var count int64
	err := db.DB.Model(&models.User{}).Where("id = ?",userID).Count(&count).Error
	if err != nil{
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

func AddNewUser(user models.User)error{
	return db.DB.Create(user).Error
}