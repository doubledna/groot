package user

import (
	"groot/internal"
	userModel "groot/internal/models/user"
)

// CreateUser create user method
func CreateUser(user userModel.User) ([]userModel.User, error) {
	var userLists []userModel.User
	result := internal.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userLists = append(userLists, user)
	return userLists, nil
}
