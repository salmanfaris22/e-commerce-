package order

import (
	"github.com/gin-gonic/gin"
)

type orderHandlerImpl struct {
	serives OrderService
}

func NewOrderHnalderV1(serives OrderService) Orderhandler {
	return &orderHandlerImpl{serives: serives}
}

func (oh orderHandlerImpl) OrderItemsChckOut(ctx *gin.Context) {
	// var cart models.Cart

}
