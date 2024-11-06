package adminOrder

import (
	"gorm.io/gorm"

	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type adminOrderRepoImpl struct {
	config *config.Config
}

func NewAdminOrderrepo(config *config.Config) AdminOrderRepo {
	return &adminOrderRepoImpl{config: config}
}

func (ao adminOrderRepoImpl) GetOrderQuary() *gorm.DB {
	return ao.config.DB.Model(&models.Order{}).Preload("Addresses").Preload("Items")
}

func (ao adminOrderRepoImpl) GetAllOrder(query *gorm.DB, orders *[]models.Order) error {
	return query.Find(&orders).Error
}

func (ao adminOrderRepoImpl) GetOrderById(orders *models.Order) error {
	return ao.config.DB.Where("id=?", orders.ID).First(&orders).Preload("Addresses").Preload("Items").Error
}

func (ao adminOrderRepoImpl) OrderStatusChncge(orders *models.Order, addreses *models.Address, orderItem *models.OrderItem) error {
	err := ao.config.DB.Model(&models.Order{}).Where("id = ?", orders.ID).Updates(orders).Error
	if err != nil {
		return err
	}
	err = ao.config.DB.Model(&models.Address{}).Where("order_id = ?", orders.ID).Updates(addreses).Error
	if err != nil {
		return err
	}
	err = ao.config.DB.Model(&models.OrderItem{}).Where("order_id = ?", orders.ID).Updates(orderItem).Error
	if err != nil {
		return err
	}
	return nil
}
