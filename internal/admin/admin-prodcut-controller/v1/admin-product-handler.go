package adminproduct

import "github.com/gin-gonic/gin"

type adminProductahndlerImpl struct {
	services AdminProductServices
}

func NewAdminProductHandler(services AdminProductServices) AdminProductHandler {
	return &adminProductahndlerImpl{services: services}
}

func (aph adminProductahndlerImpl) AddProduct(ctx *gin.Context) {

}
