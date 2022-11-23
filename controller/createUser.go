package controller

import (
	"instagram-service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	var user dto.UserDTO

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store in database

	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
