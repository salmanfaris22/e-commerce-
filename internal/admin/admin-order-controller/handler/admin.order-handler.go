package adminOrderHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	adminOrderInterFace "my-gin-app/internal/admin/admin-order-controller/interface"
	"my-gin-app/internal/models"
)

type adminOrderHandlerImpl struct {
	services adminOrderInterFace.AdminOrderrServices
}

func NewAdminOrdeHandler(services adminOrderInterFace.AdminOrderrServices) adminOrderInterFace.AdminOrderHandler {
	return &adminOrderHandlerImpl{services: services}
}

func (or adminOrderHandlerImpl) GetAllOrderAdmin(ctx *gin.Context) {

	status := ctx.Query("status")
	orders, err := or.services.GetAllOrderAdmin(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": orders,
	})

}
func (or adminOrderHandlerImpl) AdminOrderControll(ctx *gin.Context) {
	var orders models.Order
	err := ctx.BindJSON(&orders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = or.services.AdminOrderStatus(&orders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "orderUpdated ",
	})
}
func (or adminOrderHandlerImpl) AdminOrderByID(ctx *gin.Context) {
	order_id := ctx.Query("order_id")
	order, err := or.services.OrderGetByID(order_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": order,
	})
}
