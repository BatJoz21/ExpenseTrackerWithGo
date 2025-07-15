package routes

import (
	"example.com/expense-tracker-with-go/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Users Routes
	server.POST("/signin", signin)
	server.POST("/login", login)
	server.POST("/newAdmin", signinAdmin)

	// Expenses Routes
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.AuthenticateUsers)
	authenticatedRoute.GET("/expenses", getAllExpenses)
	authenticatedRoute.GET("/expenses/:id", getExpense)
	authenticatedRoute.POST("/expenses", createExpense)
	authenticatedRoute.PUT("expenses/:id", updateExpense)
	authenticatedRoute.DELETE("/expenses/:id", deleteExpense)

	// Admin Routes
	adminRoute := server.Group("/admin")
	adminRoute.Use(middleware.AuthenticateAdmin)
	adminRoute.GET("/get_users", getAllUsers)
	adminRoute.DELETE("/delete_user/:id", deleteUserByID)
}
