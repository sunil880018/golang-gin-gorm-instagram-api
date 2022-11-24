package controller

import (
	"instagram-service/dto"
	"instagram-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserHandler(context *gin.Context) {
	var user dto.UserDTO

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var db *gorm.DB
	services.CreateUser(user, db)

	context.JSON(http.StatusCreated, gin.H{
		"user": "",
	})
}
