package routes

import (
	"net/http"

	"example.com/expense-tracker-with-go/models"
	"github.com/gin-gonic/gin"
)

func signin(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to read user input"})
		return
	}

	err = newUser.SaveNewUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to register new user to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "signin successful"})
}

func login(context *gin.Context) {}
