package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type orderHandlerImpl struct {
	serives OrderService
}

func NewOrderHnalderV1(serives OrderService) Orderhandler {
	return &orderHandlerImpl{serives: serives}
}

func (oh orderHandlerImpl) OrderItemsChckOut(ctx *gin.Context) {
	id, _ := ctx.Get("user_Id")
	var tempOrder models.DemoOrder
	if err := ctx.ShouldBindJSON(&tempOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDStr, _ := id.(string)
	address, order, statusCode, err := oh.serives.OrderItems(userIDStr, &tempOrder)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"address": address,
		"order":   order,
		"message": "order successfully",
	})
}

func (oh orderHandlerImpl) CancellOrder(ctx *gin.Context) {
	id, _ := ctx.Get("user_Id")
	OrderID := ctx.Query("orderId")
	if OrderID == "" {
		ctx.JSON(400, gin.H{
			"message": "pleas select OrderId",
		})
		return
	}
	userIDStr, _ := id.(string)
	statusCode, str, err := oh.serives.CanncleorderSerivices(OrderID, userIDStr)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": str,
	})
}

func (oh orderHandlerImpl) GetAllOrder(ctx *gin.Context) {
	id, _ := ctx.Get("user_Id")
	userIDStr, _ := id.(string)
	order, statusCode, err := oh.serives.GetAllOrder(userIDStr)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": order,
	})

}
