package services

import (
	"fmt"
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"
)

func CreateUserPhotoDetails(photo dto.PhotoDTO, userId string) {
	var userPhotoDetailsObj models.UserPhotoDetails
	userPhotoDetailsObj.UserId = userId
	userPhotoDetailsObj.PhotoUrl = photo.PhotoUrl
	userPhotoDetailsRecord := database.Database.Create(&userPhotoDetailsObj)
	if userPhotoDetailsRecord != nil {
		panic(userPhotoDetailsRecord.Error)
	}
	fmt.Println(photo, userId)
}
