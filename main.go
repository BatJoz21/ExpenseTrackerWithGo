package main

import (
	"example.com/expense-tracker-with-go/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Run(":8080")
}
