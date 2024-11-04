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
	oh.serives.OrderItems(userIDStr, tempOrder)

}
