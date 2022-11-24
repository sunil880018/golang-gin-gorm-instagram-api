package services

import (
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"

	"github.com/google/uuid"
)

func CreateUser(user dto.UserDTO) (models.User, error) {
	var userObj models.User
	userObj.UserId = uuid.New().String()
	userObj.Name = user.Name
	userObj.Email = user.Email
	userObj.DOB = user.DOB
	err := database.Database.Create(&userObj).Error
	if err != nil {
		return models.User{}, err
	}
	return userObj, nil
}
