package order

import "github.com/gin-gonic/gin"

type Orderhandler interface {
	OrderItemsChckOut(ctx *gin.Context)
}

type OrderRepo interface {
}

type OrderService interface {
}
