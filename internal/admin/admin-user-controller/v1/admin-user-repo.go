package adminUser

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type adminUseImpl struct {
	config *config.Config
}

func NewAdminUserrepo(config *config.Config) AdminUserrepo {
	return &adminUseImpl{config: config}
}

func (aur adminUseImpl) GetAlluserRepo(users *[]models.User) error {
	return aur.config.DB.Preload("Orders").Preload("Wishlist").Preload("Addresses").Preload("Cart").Model(&models.User{}).Find(&users).Error
}
func (aur adminUseImpl) UpdateUser(updateUser map[string]interface{}, user *models.User, id string) error {
	return aur.config.DB.Model(&models.User{}).Where("id = ?", id).Updates(updateUser).First(&user).Error
}
func (aur adminUseImpl) GetUserById(user *models.User, id string) error {
	return aur.config.DB.Preload("Orders.Addresses").Preload("Wishlist").Preload("Cart.Items.Product").Where("id=?", id).First(&user).Error
}

func (aur adminUseImpl) UserSave(user *models.User) error {
	return aur.config.DB.Save(&user).Error
}
