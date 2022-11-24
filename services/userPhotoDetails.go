package services

import (
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"
)

func CreateUserPhotoDetails(photo dto.PhotoDTO, userId string) (models.UserPhotoDetails, error) {
	var userPhotoDetailsObj models.UserPhotoDetails
	userPhotoDetailsObj.UserId = userId
	userPhotoDetailsObj.PhotoUrl = photo.PhotoUrl
	err := database.Database.Create(&userPhotoDetailsObj).Error
	if err != nil {
		return models.UserPhotoDetails{}, err
	}
	return userPhotoDetailsObj, nil
}
