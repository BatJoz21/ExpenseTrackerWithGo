package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Users Routes
	server.POST("/signin", signin)
	server.POST("/login", login)

	// Expenses Routes
	server.GET("/expenses", getAllExpenses)
	server.GET("/expenses/:id", getExpense)
	server.POST("/expenses", createExpense)
	server.PUT("expenses/:id", updateExpense)
	server.DELETE("/expenses/:id", deleteExpense)
}
