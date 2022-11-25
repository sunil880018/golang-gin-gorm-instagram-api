package controller

import (
	"net/http"

	"instagram-service/enums"
	"instagram-service/services"

	"github.com/gin-gonic/gin"
)

func UserFollowingHandler(context *gin.Context) {
	userId := context.PostForm("userid")
	followerId := context.PostForm("followerid")
	if userId == "" || followerId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": enums.BAD_REQUEST})
		return
	}
	userFollowerObj, err := services.UserFollower(userId, followerId)
	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": enums.NOT_ACCEPTABLE})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"user": userFollowerObj,
	})
}
