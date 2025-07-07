package main

import (
	"example.com/expense-tracker-with-go/db"
	"example.com/expense-tracker-with-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
