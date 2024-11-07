package order

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type Orderhandler interface {
	OrderItemsChckOut(ctx *gin.Context)
	GetAllOrder(ctx *gin.Context)
	OrderChckOut(ctx *gin.Context)
	CancellOrder(ctx *gin.Context)
}

type OrderRepo interface {
	CreatOrder(order *models.Order) error
	FindUserCart(id string, cart *models.Cart) error
	FindcartItems(id uint, cartItems *[]models.CartItem) error
	CreateOrderItem(orderItems *models.OrderItem) error
	FindProduct(product *models.Product, pID uint) error
	SaveUpdateProduct(product *models.Product) error
	CanleOrderModel(id uint) error
	DeleteOrderItem(orderID uint) error
	DeleteCartItemsByCartID(cartID uint) error
	DeleteOrder(order uint) error
	CreatOrderAdress(address *models.Address) error
	GetAllOrders(userID string, orders *[]models.Order) error
	GetOrderById(OrderID string, userID string, order *models.Order) error
	GetAllOrderProduct(orderId uint, orderProduct *[]models.OrderItem) error
	SaveMyOrder(order *models.Order) error
	PaymentMetherd(paymen *models.Payment) error
}

type OrderService interface {
	CheckOutOrdersfromcart(id string, address models.GetOrderdetils) error
	OrderItems(id string, tempOrder *models.DemoOrder) (models.Address, models.Order, int, error)
	GetAllOrder(id string) ([]models.Order, int, error)
	CanncleorderSerivices(OrderID string, userID string) (int, string, error)
}
