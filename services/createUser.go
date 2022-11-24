package services

import (
	"fmt"
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"

	"github.com/google/uuid"
)

func CreateUser(user dto.UserDTO) {
	var userObj models.User
	userObj.UserId = uuid.New().String()
	userObj.Name = user.Name
	userObj.Email = user.Email
	userObj.DOB = user.DOB
	userRecord := database.Database.Create(&userObj)
	if userRecord != nil {
		panic(userRecord.Error)
	}
	fmt.Println(userObj)
}
