package controller

import (
	"instagram-service/dto"
	"instagram-service/enums"
	"instagram-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	var user dto.UserDTO

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": enums.BAD_REQUEST})
		return
	}

	services.CreateUser(user)

	context.JSON(http.StatusCreated, gin.H{
		"user": enums.CREATED,
	})
}
