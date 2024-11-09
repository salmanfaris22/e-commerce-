package adminproduct

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type AdminProductRepo interface {
	AddProduct(product *models.Product) error
	UpdateProdutcs(product *models.Product) error
	DeleteProductRepo(id string) error
	FindProduct(id string, existingProduct *models.Product) error
	FindImges(id uint, existingIMG *models.ProductImage) error
	UpdateImges(existingIMG *models.ProductImage) error
	SaveIMg(existingIMG *models.ProductImage) error
	FindAllImages(id uint, existingIMG *[]models.ProductImage) error
	DeleteImaged(id uint) error
}
type AdminProductServices interface {
	AddProduct(product models.Product) error
	UpdateProduct(product models.Product, id string) error
	DeleteProduct(id string) error
}
type AdminProductHandler interface {
	AddProduct(ctx *gin.Context)
	EditProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}
