package cartInterface

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type CartHandler interface {
	AddToCarthancler(ctx *gin.Context)
	GetCartItemsHandler(ctx *gin.Context)
}
type CartServices interface {
	GetAllCartItems(userID string) (int, string, error, []models.CartItem)
	AddToCartService(productId string, userId any, method string, qty int) (int, string, error)
}

type CartRepo interface {
	GetProductModelById(product *models.Product, productID string) error
	FindUserCart(userID string, cart *models.Cart) error
	CartItemfind(cartId uint, productId uint, cartItem *models.CartItem) error
	CreatCartItem(cartItem *models.CartItem) error
	DeleteCartItem(cartItem *models.CartItem) error
	Updatecart(cartItem *models.CartItem) error
	GetAllCartItems(cartItems *[]models.CartItem, cartId uint) error
	CreatCart(cart *models.Cart) error
}
