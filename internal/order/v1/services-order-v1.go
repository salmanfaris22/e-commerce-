package order

import (
	"errors"
	"strconv"
	"time"

	"my-gin-app/internal/models"
)

type orderSerivesimpl struct {
	repo OrderRepo
}

func NewOrderSerivesV1(repo OrderRepo) OrderService {
	return &orderSerivesimpl{repo: repo}
}

func (os orderSerivesimpl) OrderItems(id string, tempOrder models.DemoOrder) error {

	var total float64
	for _, t := range tempOrder.OrderItem {
		qty := float64(t.Quantity)
		productTotal := qty * t.Price
		total += productTotal
	}
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	var order models.Order
	order.TotalPrice = total
	order.CreatedAt = time.Now()
	order.Status = "pending"
	order.UserID = uint(userID)
	os.repo.CreatOrder(&order)
	var address models.Address
	address.UserID = uint(userID)
	address.CreatedAt = time.Now()
	address.City = tempOrder.City
	address.Country = tempOrder.Country
	address.OrderID = order.ID
	address.State = tempOrder.State
	address.ZipCode = tempOrder.ZipCode
	address.Street = tempOrder.Street
	for _, item := range tempOrder.OrderItem {
		var orderItems models.OrderItem
		orderItems.OrderID = order.ID
		orderItems.Quantity = item.Quantity
		orderItems.Price = item.Price
		orderItems.ProductID = item.ProductID

		err = os.repo.CreateOrderItem(&orderItems)
		if err != nil {
			return nil
		}
		err = os.ProductController(item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (os orderSerivesimpl) ProductController(pID uint, qty int, price float64) error {
	var product models.Product
	err := os.repo.FindProduct(&product, pID)
	if err != nil {
		return err
	}
	if qty > product.Stock {
		return errors.New("your qty exceeds available stock")
	}

	product.Stock -= qty
	if product.Stock <= 0 {
		product.IsAvailable = false
	}

	err = os.repo.SaveUpdateProduct(&product)
	if err != nil {
		return err
	}
	if price != product.Price {
		return errors.New("please check the price")
	}
	return nil
}
