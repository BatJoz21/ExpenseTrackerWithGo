package routes

import (
	"net/http"

	"example.com/expense-tracker-with-go/models"
	"example.com/expense-tracker-with-go/utils"
	"github.com/gin-gonic/gin"
)

func signinAdmin(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to read admin input"})
		return
	}

	err = newUser.CreateAdmin()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to register new admin to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "signin successful"})
}

func deleteUserByID(context *gin.Context) {
	uID, err := utils.FromStringToInt64(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	user, err := models.GetUserByID(uID)
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{"message": "user not found"})
		return
	}

	err = user.RemoveUserByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user has been deleted"})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsersData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to obtain all user data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Users:": users})
}
