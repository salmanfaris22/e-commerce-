package adminauthInterface

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type UserHandler interface {
	Logine(ctx *gin.Context)
	LogOut(ctx *gin.Context)
}

type UserRepo interface {
	FindByEmail(email string, user *models.Admin) error
}

type UserService interface {
	LogineUser(user models.Admin) (string, error)
}
