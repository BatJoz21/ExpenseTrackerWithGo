package routes

import (
	"net/http"

	"example.com/expense-tracker-with-go/models"
	"example.com/expense-tracker-with-go/utils"
	"github.com/gin-gonic/gin"
)

func createExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	expense.User_ID = context.GetInt64("user_id")

	err = expense.SaveExpense()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save new expense to database"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "new expenses successfully stored", "expense": expense})
}

func updateExpense(context *gin.Context) {
	eID, err := utils.FromStringToInt64(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	var oldExpense *models.Expense
	oldExpense, err = models.GetExpensebyID(eID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch expense from database"})
		return
	}

	// Check if the current active user ID == oldExpense ID
	currentUsrID := context.GetInt64("user_id")
	if oldExpense.User_ID != currentUsrID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user unauthorized to update this expense"})
		return
	}

	var updatedExpense models.Expense
	err = context.ShouldBindJSON(&updatedExpense)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed read user's input"})
		return
	}
	updatedExpense.ID = eID
	updatedExpense.User_ID = currentUsrID

	err = updatedExpense.UpdateExpenseByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update expense data on database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "expense successfully updated", "expense": updatedExpense})
}

func getExpense(context *gin.Context) {
	eID, err := utils.FromStringToInt64(context.Param("id"))
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

	currentUsrID := context.GetInt64("user_id")
	if expense.User_ID != currentUsrID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user unauthorized to fetch data from database"})
		return
	}

	context.JSON(http.StatusOK, expense)
}

func getAllExpenses(context *gin.Context) {
	currentUsrID := context.GetInt64("user_id")

	expenses, err := models.GetAllExpensesByUserID(currentUsrID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from server"})
		return
	}

	context.JSON(http.StatusOK, expenses)
}

func deleteExpense(context *gin.Context) {
	eID, err := utils.FromStringToInt64(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	expense, err := models.GetExpensebyID(eID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch data from database"})
		return
	}

	// checking current active user's ID == expense ID

	err = expense.DeleteExpenseByID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete data from database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "expense has been deleted"})
}
