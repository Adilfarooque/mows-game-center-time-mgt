package services

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
