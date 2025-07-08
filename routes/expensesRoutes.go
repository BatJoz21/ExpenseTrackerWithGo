package routes

import (
	"net/http"
	"strconv"

	"example.com/expense-tracker-with-go/models"
	"github.com/gin-gonic/gin"
)

func createExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	var userID int64 = 1 // temporary user ID
	expense.User_ID = userID

	err = expense.SaveExpense()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save new expense to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "new expenses successfully stored", "expense": expense})
}

func getExpense(context *gin.Context) {
	eID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	var expense *models.Expense
	expense, err = models.GetExpensebyID(eID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from database"})
		return
	}

	context.JSON(http.StatusOK, expense)
}

func getAllExpenses(context *gin.Context) {
	expenses, err := models.GetAllExpenses()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from server"})
		return
	}

	context.JSON(http.StatusOK, expenses)
}
