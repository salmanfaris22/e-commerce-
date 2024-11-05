package adminauth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type UserHandlerImpl struct {
	userService UserService
}

func NewAdminUserHandlerV1(service UserService) UserHandler {
	return &UserHandlerImpl{userService: service}
}

func (uh *UserHandlerImpl) Logine(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.BindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}

	str, err := uh.userService.LogineUser(admin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user information", "details": err.Error()})
		return
	}

	ctx.SetCookie("adminToken", str, int(7*24*time.Hour), "/", "localhost", true, true)

	ctx.JSON(200, gin.H{
		"message": "Logine successful",
	})
}

func (uh *UserHandlerImpl) LogOut(ctx *gin.Context) {
	ctx.SetCookie("adminToken", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}
