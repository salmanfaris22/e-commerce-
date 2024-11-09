package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-gin-app/internal/models"
)

type ProductHandle interface {
	GetAllProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	SerchProduct(ctx *gin.Context)
	FilterProducts(ctx *gin.Context)
}

type ProductRepo interface {
	GetAllProductModel(product *[]models.Product) error
	GetProductModelById(product *models.Product) error
	SearchProductRepo(product *[]models.Product, searchItem string) error
	FilterQuery() *gorm.DB
	QueryFindProduct(query *gorm.DB, products *[]models.Product) error
}

type ProducctServices interface {
	GetAllProduct() ([]models.Product, error)
	GetIDProductService(productID string) (models.Product, error)
	SerchProductService(searchItem string) ([]models.Product, error)
	FilterProduct(filter models.Filter, Available, maxPriceStr, minPriceStr string) ([]models.Product, error)
}
