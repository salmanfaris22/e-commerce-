package adminauthRepo

import (
	"my-gin-app/config"
	adminauthInterface "my-gin-app/internal/admin/auth/interface"
	"my-gin-app/internal/models"
)

type UserRepoImpl struct {
	config *config.Config
}

func NewAdminUserRepoV1(config *config.Config) adminauthInterface.UserRepo {
	return &UserRepoImpl{config: config}
}

func (u UserRepoImpl) FindByEmail(email string, user *models.Admin) error {
	return u.config.DB.Where("email = ?", email).First(user).Error
}
