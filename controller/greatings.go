package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greetings(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"mesg": "Hello",
	})
}
