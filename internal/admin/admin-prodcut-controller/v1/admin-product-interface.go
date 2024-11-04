package adminproduct

import "github.com/gin-gonic/gin"

type AdminProductRepo interface {
}
type AdminProductServices interface {
}
type AdminProductHandler interface {
	AddProduct(ctx *gin.Context)
}
