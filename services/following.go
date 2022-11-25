package services

import (
	"instagram-service/database"
	"instagram-service/models"
)

func UserFollower(userId string, follower string) (models.UserFollower, error) {
	var userFollower models.UserFollower
	userFollower.UserId = userId
	userFollower.Followers = follower
	err := database.Database.Create(&userFollower).Error
	if err != nil {
		return models.UserFollower{}, err
	}
	return userFollower, nil
}
