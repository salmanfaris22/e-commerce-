package adminUserInterface

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type AdminUserServices interface {
	GetAlluser() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	EditUser(updateUser map[string]interface{}, id string) (models.User, error)
	BlockuserServieces(UserStatus string, user_id string) (string, error)
}
type AdminUserrepo interface {
	GetAlluserRepo(users *[]models.User) error
	UpdateUser(updateUser map[string]interface{}, user *models.User, id string) error
	GetUserById(user *models.User, id string) error
	UserSave(user *models.User) error
}
type AdminUserHandler interface {
	GetAllUSer(ctx *gin.Context)
	EditUser(ctx *gin.Context)
	GetUSer(ctx *gin.Context)
	BlockUser(ctx *gin.Context)
}
