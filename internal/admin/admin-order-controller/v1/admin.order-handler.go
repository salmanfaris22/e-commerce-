package adminOrder

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type AdminOrderHandler interface {
	GetAllOrderAdmin(ctx *gin.Context)
	AdminOrderControll(ctx *gin.Context)
}

type adminOrderHandlerImpl struct {
	services AdminOrderrServices
}

func NewAdminOrdeHandler(services AdminOrderrServices) AdminOrderHandler {
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
		"message": orders,
	})
}
