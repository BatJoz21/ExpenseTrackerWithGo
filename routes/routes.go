package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Expenses Routes
	server.GET("/expenses", getAllExpenses)
	server.POST("expenses", createExpense)
}
