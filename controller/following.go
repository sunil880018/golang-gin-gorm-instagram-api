package controller

import (
	"net/http"

	"instagram-service/dto"

	"github.com/gin-gonic/gin"
)

func UserFollowingHandler(context *gin.Context) {
	var user dto.UserDTO
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"user": user.UserId,
	})
}
