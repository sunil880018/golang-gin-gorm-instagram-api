package controller

import (
	"net/http"

	"instagram-service/dto"
	"instagram-service/enums"

	"github.com/gin-gonic/gin"
)

func UserFollowingHandler(context *gin.Context) {
	var user dto.UserDTO
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": enums.BAD_REQUEST})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"user": user.UserId,
	})
}
