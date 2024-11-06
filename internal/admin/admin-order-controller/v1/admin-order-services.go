package adminOrder

import "my-gin-app/internal/models"

type AdminOrderrServices interface {
	AdminOrderStatus(orders *models.Order) error
	GetAllOrderAdmin(status string) ([]models.Order, error)
}

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
	err := as.repo.GetOrderById(orders)
	if err != nil {
		return err
	}
	addresses := orders.Addresses[0]
	order_item := orders.Items[0]
	err = as.repo.OrderStatusChncge(orders, &addresses, &order_item)
	if err != nil {
		return err
	}
	return nil
}
