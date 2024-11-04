package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Logine(ctx *gin.Context)
}

type UserHandlerImpl struct {
	userService UserService
}

func NewUserHandlerV1(service UserService) UserHandler {
	return &UserHandlerImpl{userService: service}
}

func (uh *UserHandlerImpl) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}
	newUser, accessToken, refreshToken, err := uh.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+accessToken)
	ctx.SetCookie("refreshToken", refreshToken, int(7*24*time.Hour), "/", "localhost", true, true)

	ctx.JSON(200, gin.H{
		"message":     "register successful",
		"accessToken": accessToken,
		"name":        newUser.FirstName,
	})
}

func (uh *UserHandlerImpl) Logine(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}

	newUser, accessToken, refreshToken, err := uh.userService.LogineUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+accessToken)
	ctx.SetCookie("refreshToken", refreshToken, int(7*24*time.Hour), "/", "localhost", true, true)

	ctx.JSON(200, gin.H{
		"message":     "Logine successful",
		"accessToken": accessToken,
		"name":        newUser.FirstName,
	})
}