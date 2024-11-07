package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"my-gin-app/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		var accessToken string
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				accessToken = parts[1]
			}
		}

		if accessToken != "" {
			valid, err := utils.ValidateToken(accessToken)
			if err == nil && valid {
				id, err := GetUserId(ctx)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": "pleas Logine"})
					ctx.Abort()
					return
				}
				ctx.Set("user_Id", id)
				ctx.Next()
				return
			}
		}

		refreshToken, err := ctx.Cookie("refreshToken")
		if err != nil {
			fmt.Println(err)
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

		newAccessToken, err := utils.GenerateAccessTokenFromRefreshToken(refreshToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to generate access token"})
			ctx.Abort()
			return
		}

		ctx.Header("Authorization", "Bearer "+newAccessToken)

		ctx.SetCookie("refreshToken", refreshToken, int(7*24*time.Hour.Seconds()), "/", "localhost", true, true)
		id, err := GetUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "pleas Logine"})
			ctx.Abort()
			return
		}

		ctx.Set("user_Id", id)
		ctx.Next()
	}
}

func GetUserId(ctx *gin.Context) (string, error) {
	token, err := ctx.Cookie("refreshToken")
	if err != nil {
		return "", err
	}
	id, err := utils.DecodeRefreshToken(token)
	if err != nil {
		return "", err
	}
	userIDStr := strconv.FormatUint(uint64(id), 10)
	return userIDStr, nil
}
