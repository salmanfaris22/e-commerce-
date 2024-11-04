package cart

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type cartRepoImpl struct {
	config config.Config
}

func NewrepoCartV1(config *config.Config) CartRepo {
	return &cartRepoImpl{config: *config}
}
func (cr cartRepoImpl) GetAllCartItems(cartItems *[]models.CartItem, cartId uint) error {
	return cr.config.DB.Preload("Product").Where("cart_id=?", cartId).Find(&cartItems).Error
}
func (cr cartRepoImpl) GetProductModelById(product *models.Product, productID string) error {
	return cr.config.DB.Preload("Images").Where("id=?", productID).First(&product).Error
}

func (cr cartRepoImpl) Updatecart(cartItem *models.CartItem) error {
	return cr.config.DB.Save(&cartItem).Error
}
func (cr cartRepoImpl) DeleteCartItem(cartItem *models.CartItem) error {
	return cr.config.DB.Delete(&cartItem).Error
}

func (cr cartRepoImpl) CartItemfind(cartId uint, productId uint, cartItem *models.CartItem) error {
	return cr.config.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&cartItem).Error
}

func (cr cartRepoImpl) CreatCartItem(qty int, cart models.Cart, productId uint, cartItem *models.CartItem) error {
	cartItem.Quantity = qty
	cartItem.ProductID = productId
	cartItem.CartID = cart.ID
	return cr.config.DB.Create(&cartItem).Error
}

func (cr cartRepoImpl) FindUserCart(userID string, cart *models.Cart) error {
	return cr.config.DB.Where("user_id=?", userID).First(&cart).Error

}

func (cr cartRepoImpl) CreatCart(cart *models.Cart, id uint) error {

	cart.UserID = id
	return cr.config.DB.Create(&cart).Error
}
