package services

import (
	"errors"
	"mows-game-center-time-mgt/models"
	"mows-game-center-time-mgt/repository"
)

// import (
// 	"errors"
// 	"mows-game-center-time-mgt/utils/models"
// )

// var users []models.User

// // register new user
// func RegisterUser(newUser models.User) models.User {
// 	newUser.ID = len(users) + 1
// 	users = append(users, newUser)
// 	return newUser
// }

// // Update user details
// func UpdateUser(id int, updatUser models.User) (models.User, error) {
// 	for i, user := range users {
// 		if user.ID == id {
// 			users[i] = updatUser
// 			return users[i], nil
// 		}
// 	}
// 	return models.User{}, errors.New("user not found")
// }

// // Remove a user
// func RemoveUser(id int) error {
// 	for i, user := range users {
// 		if user.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return errors.New("user not found")
// }

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

// DeleteUser by ID using repository
func DeleteUser(id int) error {
	//Check the user if exists to attempt to delete
	userExists := repository.CheckUserAvailabilityWithID(id)
	if !userExists {
		return errors.New("user not found")
	}
	//If user exists,
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

// Fetch use by ID
func GetUserByID(id int) (models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func AddNewUser(user *models.User)error{
	return repository.AddNewUser(*user)
}