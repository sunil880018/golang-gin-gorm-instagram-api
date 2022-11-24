package services

import (
	"fmt"
	"instagram-service/dto"
	"instagram-service/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(user dto.UserDTO, db *gorm.DB) {
	var userObj models.User
	userObj.UserId = uuid.New().String()
	userObj.Name = user.Name
	userObj.Email = user.Email
	userObj.DOB = user.DOB
	result := db.Create(&userObj)
	fmt.Println(result)
}
