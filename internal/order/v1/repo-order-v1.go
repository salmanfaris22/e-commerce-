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
func (or orderRepoImpl) CreatOrderAdress(address *models.Address) error {
	return or.config.DB.Create(&address).Error
}

func (or orderRepoImpl) GetAllOrders(userID string, orders *[]models.Order) error {
	return or.config.DB.Preload("Addresses").Preload("Items").Where("user_id = ? AND status!=?", userID, "canceled").Find(&orders).Error
}
func (or orderRepoImpl) GetOrderById(OrderID string, userID string, order *models.Order) error {
	return or.config.DB.Where("id=? AND user_id=?", OrderID, userID).First(&order).Error
}

func (or orderRepoImpl) GetAllOrderProduct(orderId uint, orderProduct *[]models.OrderItem) error {
	return or.config.DB.Where("order_id=?", orderId).Find(&orderProduct).Error
}
func (or orderRepoImpl) SaveMyOrder(order *models.Order) error {
	return or.config.DB.Save(&order).Error
}
