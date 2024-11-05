package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-gin-app/pkg/utils"
)

func AdminMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken, err := ctx.Cookie("adminToken")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in"})
			ctx.Abort()
			return
		}
		valid, err := utils.ValidateToken(refreshToken)
		if err != nil || !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in"})
			ctx.Abort()
			return
		}
		if valid {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in"})
		ctx.Abort()
		return
	}
}
