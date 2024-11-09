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
	return ao.config.DB.Model(&models.Order{}).Preload("Addresses").Preload("Items.Product")
}

func (ao adminOrderRepoImpl) GetAllOrder(query *gorm.DB, orders *[]models.Order) error {
	return query.Preload("Addresses").Preload("Items.Product.Images").Find(&orders).Error
}

func (ao adminOrderRepoImpl) GetOrderById(orders *models.Order) error {
	return ao.config.DB.Preload("Addresses").Preload("Items.Product.Images").Where("id=?", orders.ID).First(&orders).Error
}

func (or adminOrderRepoImpl) UpdateOrder(order *models.Order) error {

	err := or.config.DB.Model(&models.Order{}).Where("id = ?", order.ID).Updates(order).Error
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err = or.config.DB.Model(&models.OrderItem{}).Where("id = ?", item.ID).Updates(item).Error
		if err != nil {
			return err
		}
	}
	return nil
}
