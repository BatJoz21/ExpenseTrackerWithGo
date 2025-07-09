package routes

import (
	"example.com/expense-tracker-with-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Users Routes
	server.POST("/signin", signin)
	server.POST("/login", login)

	// Expenses Routes
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)
	authenticatedRoute.GET("/expenses", getAllExpenses)
	authenticatedRoute.GET("/expenses/:id", getExpense)
	authenticatedRoute.POST("/expenses", createExpense)
	authenticatedRoute.PUT("expenses/:id", updateExpense)
	authenticatedRoute.DELETE("/expenses/:id", deleteExpense)
}
