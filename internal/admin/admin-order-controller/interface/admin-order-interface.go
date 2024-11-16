package adminOrderInterFace

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-gin-app/internal/models"
)

type AdminOrderRepo interface {
	GetOrderQuary() *gorm.DB
	GetAllOrder(query *gorm.DB, orders *[]models.Order) error
	GetOrderById(orders *models.Order) error
	UpdateOrder(order *models.Order) error
}
type AdminOrderrServices interface {
	AdminOrderStatus(orders *models.Order) error
	GetAllOrderAdmin(status string) ([]models.Order, error)
	OrderGetByID(order_id string) (models.Order, error)
}
type AdminOrderHandler interface {
	GetAllOrderAdmin(ctx *gin.Context)
	AdminOrderControll(ctx *gin.Context)
	AdminOrderByID(ctx *gin.Context)
}
