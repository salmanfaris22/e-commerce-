package cart

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type cartHavndelerImpl struct {
	cartServices CartServices
}

func NewHandleCartV1(services CartServices) CartHandler {
	return &cartHavndelerImpl{cartServices: services}
}

func (ch cartHavndelerImpl) GetCartItemsHandler(ctx *gin.Context) {
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)

	status, str, err, items := ch.cartServices.GetAllCartItems(userIDStr)
	if err != nil {
		ctx.JSON(status, gin.H{
			"error":   err,
			"message": "no cart items..!",
		})
		return
	}
	ctx.JSON(status, gin.H{
		"message": str,
		"items":   items,
	})

}

func (ch cartHavndelerImpl) AddToCarthancler(ctx *gin.Context) {
	productId := ctx.Query("productId")
	method := ctx.Query("use")

	id, _ := ctx.Get("user_Id")

	qty, errs := strconv.Atoi(ctx.Query("qty"))
	if errs != nil {
		qty = 1
	}
	status, str, errs := ch.cartServices.AddToCartService(productId, id, method, qty)
	if errs != nil {
		ctx.JSON(status, gin.H{
			"error":   errs,
			"message": "can't add to cart..!",
		})
		return
	}
	ctx.JSON(status, gin.H{
		"message": str,
		"id":      productId,
		"userID":  id,
	})
}
