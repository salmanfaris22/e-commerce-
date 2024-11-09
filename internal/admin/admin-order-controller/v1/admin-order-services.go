package adminOrder

import (
	"strconv"

	"my-gin-app/internal/models"
)

type adminOrderServicesoImpl struct {
	repo AdminOrderRepo
}

func NewAdminOrdeServices(repo AdminOrderRepo) AdminOrderrServices {
	return &adminOrderServicesoImpl{repo: repo}
}

func (as adminOrderServicesoImpl) GetAllOrderAdmin(status string) ([]models.Order, error) {
	var orders []models.Order
	query := as.repo.GetOrderQuary()
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := as.repo.GetAllOrder(query, &orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
func (as adminOrderServicesoImpl) AdminOrderStatus(orders *models.Order) error {
	var demo models.Order
	demo.ID = orders.ID
	err := as.repo.GetOrderById(&demo)
	if err != nil {
		return err
	}
	err = as.repo.UpdateOrder(orders)
	if err != nil {
		return err
	}
	return nil
}

func (as adminOrderServicesoImpl) OrderGetByID(order_id string) (models.Order, error) {
	var orders models.Order
	num, err := strconv.ParseUint(order_id, 10, 64)
	if err != nil {
		return orders, err
	}
	orders.ID = uint(num)
	err = as.repo.GetOrderById(&orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
