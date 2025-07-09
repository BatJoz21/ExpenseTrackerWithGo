package routes

import (
	"net/http"

	"example.com/expense-tracker-with-go/models"
	"example.com/expense-tracker-with-go/utils"
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

func login(context *gin.Context) {
	var loggedUser models.User
	err := context.ShouldBindJSON(&loggedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to read user input"})
		return
	}

	err = loggedUser.ValidatingCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(loggedUser.ID, loggedUser.Email, loggedUser.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull", "token": token})
}
