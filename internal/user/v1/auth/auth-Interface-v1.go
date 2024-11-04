package auth

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Logine(ctx *gin.Context)
}

type UserRepo interface {
	Save(user *models.User) error
	TokenSave(token *models.UserToken) error
	FindByEmail(email string, user *models.User) error
	FindRefreshToken(token *models.UserToken, id uint) error
}

type UserService interface {
	CreateUser(user models.User) (models.User, string, string, error)
	LogineUser(user models.User) (models.User, string, string, error)
}
