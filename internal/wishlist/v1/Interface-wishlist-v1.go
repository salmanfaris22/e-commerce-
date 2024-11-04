package wishlist

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type WishListRepo interface {
	GetPpduct(product *models.Product, productId string) error
	FindWishlist(userID string, wishlist *models.Wishlist) error
	CreatWishlist(wishlist *models.Wishlist) error
	FindWishlistItem(wishlistId uint, productId uint, listItem *models.WishlistItem) error
	CreatWishlistItem(listItem *models.WishlistItem) error
	DeleteWishlistItem(listItem *models.WishlistItem) error
}

type WishListServices interface {
	WishListAddremove(userID, productID string) (string, error)
}

type WishListHandler interface {
	WishListController(gin *gin.Context)
}
