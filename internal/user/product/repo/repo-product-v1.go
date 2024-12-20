package productRepo

import (
	"gorm.io/gorm"

	"my-gin-app/config"
	"my-gin-app/internal/models"
	productInterface "my-gin-app/internal/user/product/interface"
)

type producctRepoImpl struct {
	config *config.Config
}

func NewProducctRepoV1(config *config.Config) productInterface.ProductRepo {
	return &producctRepoImpl{config: config}
}

func (pr producctRepoImpl) GetAllProductModel(product *[]models.Product) error {
	return pr.config.DB.Preload("Images").Model(&models.Product{}).Find(&product).Error
}

func (pr producctRepoImpl) GetProductModelById(product *models.Product) error {
	return pr.config.DB.Preload("Images").Preload("Reviews").Where("id=?", product.ID).First(&product).Error
}
func (pr producctRepoImpl) SearchProductRepo(product *[]models.Product, searchItem string) error {
	return pr.config.DB.Preload("Images").Where("name ILIKE ? OR description ILIKE ? OR category ILIKE ? OR brand ILIKE ?", "%"+searchItem+"%", "%"+searchItem+"%", "%"+searchItem+"%", "%"+searchItem+"%").Find(&product).Error
}
func (pr producctRepoImpl) FilterQuery() *gorm.DB {
	return pr.config.DB.Preload("Images").Model(&models.Product{})
}
func (pr producctRepoImpl) QueryFindProduct(query *gorm.DB, products *[]models.Product) error {
	return query.Find(&products).Error
}
