package order

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type Orderhandler interface {
	OrderItemsChckOut(ctx *gin.Context)
}

type OrderRepo interface {
	CreatOrder(order *models.Order) error
	CreateOrderItem(orderItems *models.OrderItem) error
	FindProduct(product *models.Product, pID uint) error
	SaveUpdateProduct(product *models.Product) error
	DeleteOrderItem(orderID uint) error
	DeleteOrder(order uint) error
}

type OrderService interface {
	OrderItems(id string, order models.DemoOrder) error
}
