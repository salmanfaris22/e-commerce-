package authHandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
	authInterface "my-gin-app/internal/user/auth/interface"
)

type UserHandlerImpl struct {
	userService authInterface.UserService
}

func NewUserHandlerV1(service authInterface.UserService) authInterface.UserHandler {
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
		"user_id":     newUser.ID,
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
		"user_id":     newUser.ID,
	})
}
func (uh *UserHandlerImpl) LogOut(ctx *gin.Context) {
	ctx.Header("Authorization", "")
	ctx.SetCookie("userId", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refreshToken", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}
