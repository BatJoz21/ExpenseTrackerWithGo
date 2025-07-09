package middleware

import (
	"net/http"

	"example.com/expense-tracker-with-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	user_id, user_role, err := utils.VerifiedToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	context.Set("user_id", user_id)
	context.Set("user_role", user_role)
	context.Next()
}
