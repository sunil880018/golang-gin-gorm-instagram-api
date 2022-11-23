package controller

import (
	"instagram-service/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	var user schemas.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// store in database

	context.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
