package wishlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type wishListHandlerImpl struct {
	servies WishListServices
}

func NewWishListHanlder(servies WishListServices) WishListHandler {
	return &wishListHandlerImpl{servies: servies}
}

func (wh wishListHandlerImpl) WishListController(ctx *gin.Context) {
	productId := ctx.Query("productId")
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)
	str, err := wh.servies.WishListAddremove(productId, userIDStr)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error(), "messgae": str})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": str,
	})

}
func (wh wishListHandlerImpl) GetAllwishlistItem(ctx *gin.Context) {

	id, _ := ctx.Get("user_Id")

	userIDStr, _ := id.(string)
	str, product, err := wh.servies.GetAllWihslistItems(userIDStr)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error(), "messgae": str})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": product,
	})

}
