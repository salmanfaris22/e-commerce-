package adminproduct

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type AdminProductRepo interface {
	AddProduct(product *models.Product) error
	UpdateProdutcs(updates interface{}, id string) error
	DeleteProductRepo(id string) error
}
type AdminProductServices interface {
	AddProduct(product models.Product) error
	UpdateProduct(updates interface{}, id string) error
	DeleteProduct(id string) error
}
type AdminProductHandler interface {
	AddProduct(ctx *gin.Context)
	EditProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}
