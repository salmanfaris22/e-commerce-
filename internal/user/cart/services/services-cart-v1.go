package cartServices

import (
	"errors"
	"fmt"
	"net/http" // Import net/http for HTTP status codes
	"strconv"

	"gorm.io/gorm"

	"my-gin-app/internal/models"
	cartInterface "my-gin-app/internal/user/cart/interface"
)

type cartServiceImpl struct {
	cartrepo cartInterface.CartRepo
}

func NewServiceCartV1(repo cartInterface.CartRepo) cartInterface.CartServices {
	return &cartServiceImpl{cartrepo: repo}
}

func (cs cartServiceImpl) GetAllCartItems(userID string) (int, string, error, []models.CartItem) {
	var cart models.Cart
	var cartItems []models.CartItem
	err := cs.cartrepo.FindUserCart(userID, &cart)
	if err != nil {
		return http.StatusInternalServerError, "", err, cartItems
	}

	err = cs.cartrepo.GetAllCartItems(&cartItems, cart.ID)
	if err != nil {
		return http.StatusInternalServerError, "", err, cartItems
	}

	return 200, "successfully", nil, cartItems
}

func (cs cartServiceImpl) AddToCartService(productId string, id any, method string, qty int) (int, string, error) {
	var product models.Product

	userIDStr, ok := id.(string)
	if !ok {
		return http.StatusBadRequest, "", errors.New("invalid user ID format; expected string")
	}

	err := cs.cartrepo.GetProductModelById(&product, productId)
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	if product.Stock == 0 || !product.IsAvailable {
		return http.StatusNotFound, "", errors.New("out of stock")
	}

	var cart models.Cart
	err = cs.cartrepo.FindUserCart(userIDStr, &cart)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			conId, _ := strconv.Atoi(userIDStr)
			cart.UserID = uint(conId)
			err = cs.cartrepo.CreatCart(&cart)
			if err != nil {
				return http.StatusInternalServerError, "", err
			}
		} else {
			return http.StatusInternalServerError, "", err
		}
		return http.StatusInternalServerError, "", err
	}

	var cartItem models.CartItem
	err = cs.cartrepo.CartItemfind(cart.ID, product.ID, &cartItem)
	if method == "remove" {
		cartItem.Quantity = 0
		qty = 0
	}

	if err != nil {

		if err == gorm.ErrRecordNotFound && qty > 0 {

			cartItem.Quantity = qty
			cartItem.ProductID = product.ID
			cartItem.CartID = cart.ID

			err := cs.cartrepo.CreatCartItem(&cartItem)
			if err != nil {
				return http.StatusInternalServerError, "", errors.New("cart item creation error")
			}
			return http.StatusOK, "Cart item added successfully", nil
		} else {
			fmt.Println(err)
			return http.StatusInternalServerError, err.Error(), err
		}
	} else {

		if qty <= 0 {
			cartItem.Quantity += qty
			if cartItem.Quantity <= 0 {
				err = cs.cartrepo.DeleteCartItem(&cartItem)
				if err != nil {
					return http.StatusInternalServerError, "", errors.New("Unable to remove item from cart")
				}
				return http.StatusOK, "Cart item removed successfully", nil
			}
		} else {
			cartItem.Quantity += qty
		}

		err = cs.cartrepo.Updatecart(&cartItem)
		if err != nil {
			return http.StatusInternalServerError, "", errors.New("can't save quantity")
		}
		if method == "remove" {
			return http.StatusOK, "Cart item removed successfully", nil
		}
		return http.StatusOK, "Cart quantity increased to " + strconv.Itoa(cartItem.Quantity), nil
	}
}
