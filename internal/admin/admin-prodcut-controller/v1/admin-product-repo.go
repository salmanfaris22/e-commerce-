package adminproduct

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type adminProductREpoImpl struct {
	config *config.Config
}

func NewAdminProductReposetries(config *config.Config) AdminProductRepo {
	return &adminProductREpoImpl{config: config}
}

func (apr adminProductREpoImpl) AddProduct(product *models.Product) error {
	return apr.config.DB.Create(&product).Error
}

func (apr adminProductREpoImpl) UpdateProdutcs(updates interface{}, id string) error {
	return apr.config.DB.Model(&models.Product{}).Where("id=?", id).Updates(updates).Error
}
func (apr adminProductREpoImpl) DeleteProductRepo(id string) error {
	return apr.config.DB.Delete(&models.Product{}, id).Error
}
