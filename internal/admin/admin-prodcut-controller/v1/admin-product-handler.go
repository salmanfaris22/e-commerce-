package adminproduct

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type adminProductahndlerImpl struct {
	services AdminProductServices
}

func NewAdminProductHandler(services AdminProductServices) AdminProductHandler {
	return &adminProductahndlerImpl{services: services}
}

func (aph adminProductahndlerImpl) AddProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Product",
		})
		return
	}
	err = aph.services.AddProduct(product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Product",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "product successful added",
	})

}
func (aph adminProductahndlerImpl) EditProduct(ctx *gin.Context) {
	id := ctx.Query("product_id")
	updates := make(map[string]interface{})
	err := ctx.BindJSON(&updates)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Product",
			"err":     err.Error(),
		})
		return
	}
	err = aph.services.UpdateProduct(updates, id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Product",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "product successful Updated",
	})

}
func (aph adminProductahndlerImpl) DeleteProduct(ctx *gin.Context) {
	id := ctx.Query("product_id")
	err := aph.services.DeleteProduct(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "cant delete Product",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "product successful DELETED",
	})

}
