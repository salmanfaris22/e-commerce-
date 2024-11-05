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
	return apr.config.DB.Preload("Images").Create(&product).Error
}

func (apr adminProductREpoImpl) FindProduct(id string, existingProduct *models.Product) error {
	return apr.config.DB.Preload("Images").First(&existingProduct, id).Error
}

func (apr adminProductREpoImpl) UpdateProdutcs(existingProduct *models.Product) error {
	return apr.config.DB.Preload("Images").Save(&existingProduct).Error
}
func (apr adminProductREpoImpl) DeleteProductRepo(id string) error {
	return apr.config.DB.Preload("Images").Delete(&models.Product{}, id).Error
}

func (apr adminProductREpoImpl) FindImges(id uint, existingIMG *models.ProductImage) error {
	return apr.config.DB.First(&existingIMG, id).Error
}
func (apr adminProductREpoImpl) UpdateImges(existingIMG *models.ProductImage) error {
	return apr.config.DB.Save(&existingIMG).Error
}

func (apr adminProductREpoImpl) SaveIMg(existingIMG *models.ProductImage) error {
	return apr.config.DB.Create(&existingIMG).Error
}
func (apr adminProductREpoImpl) FindAllImages(id uint, existingIMG *[]models.ProductImage) error {
	return apr.config.DB.Where("product_id=?", id).Find(&existingIMG).Error
}
func (apr adminProductREpoImpl) DeleteImaged(id uint) error {
	return apr.config.DB.Delete(&models.ProductImage{}, id).Error
}
