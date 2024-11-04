package order

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type orderRepoImpl struct {
	config config.Config
}

func NewOrderRepoV1(config config.Config) OrderRepo {
	return &orderRepoImpl{config: config}
}

func (or orderRepoImpl) CreatOrder(order *models.Order) error {
	return or.config.DB.Create(&order).Error
}

func (or orderRepoImpl) CreateOrderItem(orderItems *models.OrderItem) error {
	return or.config.DB.Create(&orderItems).Error
}

func (or orderRepoImpl) FindProduct(product *models.Product, pID uint) error {
	return or.config.DB.Where("id=?", pID).First(&product).Error
}
func (or orderRepoImpl) SaveUpdateProduct(product *models.Product) error {
	return or.config.DB.Save(&product).Error
}
func (or orderRepoImpl) DeleteOrderItem(orderID uint) error {
	return or.config.DB.Where("order_id = ?", orderID).Delete(&models.OrderItem{}).Error
}
func (or orderRepoImpl) DeleteOrder(order uint) error {
	return or.config.DB.Where("id = ?", order).Delete(&models.Order{}).Error
}
