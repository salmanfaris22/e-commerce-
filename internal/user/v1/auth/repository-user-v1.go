package repositories

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type UserRepo interface {
	Save(user *models.User) error
	TokenSave(token *models.UserToken) error
	FindByEmail(email string, user *models.User) error
	FindRefreshToken(token *models.UserToken, id uint) error
}

type UserRepoImpl struct {
	config *config.Config
}

func NewUserRepoV1(config *config.Config) UserRepo {
	return &UserRepoImpl{config: config}
}

func (u UserRepoImpl) Save(user *models.User) error {
	return u.config.DB.Save(&user).Error
}

func (u UserRepoImpl) FindByEmail(email string, user *models.User) error {
	return u.config.DB.Where("email = ?", email).First(user).Error
}
func (u UserRepoImpl) TokenSave(token *models.UserToken) error {
	return u.config.DB.Save(&token).Error
}

func (u UserRepoImpl) FindRefreshToken(token *models.UserToken, id uint) error {
	return u.config.DB.Where("user_id = ?", id).First(&token).Error
}
