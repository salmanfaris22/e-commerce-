package orderServices

import (
	"errors"
	"fmt"
	"net/http" // Import net/http for HTTP status codes
	"strconv"
	"time"

	"my-gin-app/internal/models"
	orderInterface "my-gin-app/internal/user/order/interface"
)

type orderServiceImpl struct {
	repo orderInterface.OrderRepo
}

func NewOrderServiceV1(repo orderInterface.OrderRepo) orderInterface.OrderService {
	return &orderServiceImpl{repo: repo}
}

func (os orderServiceImpl) CanncleorderSerivices(OrderID string, userID string) (int, string, error) {
	var order models.Order
	err := os.repo.GetOrderById(OrderID, userID, &order)
	if err != nil {
		return http.StatusInternalServerError, "can't find Order", err
	}
	order.Status = "Cancelled"
	order.UpdatedAt = time.Now()
	err = os.CancelOrderController(&order)
	if err != nil {
		return http.StatusInternalServerError, "can't find Order", err
	}

	if err != nil {
		return http.StatusInternalServerError, "can't find Order", err
	}
	err = os.repo.SaveMyOrder(&order)
	if err != nil {
		return http.StatusInternalServerError, "can't find Order", err
	}

	return 200, "order cancelled", nil

}
func (os orderServiceImpl) CancelOrderController(order *models.Order) error {
	var orderProduct []models.OrderItem
	err := os.repo.GetAllOrderProduct(order.ID, &orderProduct)
	if err != nil {
		return err
	}
	for _, item := range orderProduct {
		err = os.repo.CanleOrderModel(item.ID)
		var product models.Product
		err = os.repo.FindProduct(&product, item.ID)
		if err != nil {
			return err
		}
		product.Stock += item.Quantity
		if product.Stock >= 0 {
			product.IsAvailable = true
		}
		err = os.repo.SaveUpdateProduct(&product)
		if err != nil {
			return err
		}
	}
	return nil
}
func (os orderServiceImpl) GetAllOrder(id string) ([]models.Order, int, error) {
	var orders []models.Order
	err := os.repo.GetAllOrders(id, &orders)
	if err != nil {
		return orders, http.StatusInternalServerError, err
	}
	return orders, 200, nil
}

func (os orderServiceImpl) OrderItems(id string, tempOrder *models.DemoOrder) (models.Address, models.Order, int, error) {
	var order models.Order
	var address models.Address
	var total float64

	for _, t := range tempOrder.OrderItem {
		qty := float64(t.Quantity)
		productTotal := qty * t.Price
		total += productTotal
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return address, order, http.StatusBadRequest, err
	}

	order.TotalPrice = total
	order.CreatedAt = time.Now()
	order.Status = tempOrder.Pymentmetherd
	order.UserID = uint(userID)
	err = os.repo.CreatOrder(&order)
	if err != nil {
		return address, order, http.StatusInternalServerError, err
	}

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
		orderItems.OrderStatus = "Pending"
		err = os.repo.CreateOrderItem(&orderItems)
		if err != nil {
			return address, order, http.StatusInternalServerError, err
		}
		err = os.ProductController(item.ProductID, item.Quantity, item.Price)
		if err != nil {
			os.repo.DeleteOrderItem(order.ID)
			os.repo.DeleteOrder(order.ID)
			return address, order, http.StatusBadRequest, err
		}
	}

	err = os.repo.CreatOrderAdress(&address)
	if err != nil {
		return address, order, http.StatusInternalServerError, err
	}
	return address, order, http.StatusOK, nil
}

func (os orderServiceImpl) ProductController(pID uint, qty int, price float64) error {
	var product models.Product

	err := os.repo.FindProduct(&product, pID)

	if err != nil {
		fmt.Println("sdflmnskfj")
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

func (os orderServiceImpl) CheckOutOrdersfromcart(id string, tempOrder models.GetOrderdetils) error {
	var cart models.Cart

	err := os.repo.FindUserCart(id, &cart)
	if err != nil {
		return err
	}

	var cartItems []models.CartItem
	err = os.repo.FindcartItems(cart.ID, &cartItems)
	if len(cartItems) == 0 {
		return errors.New("not cart items")
	}
	if err != nil {
		return err
	}
	var order models.Order
	var total float64
	for _, t := range cartItems {
		qty := float64(t.Quantity)
		var tempProduct models.Product
		err = os.repo.FindProduct(&tempProduct, t.ProductID)
		if err != nil {
			return err
		}
		productTotal := qty * tempProduct.Price
		total += productTotal
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	order.TotalPrice = total
	order.CreatedAt = time.Now()
	order.Status = tempOrder.Mtherd
	order.UserID = uint(userID)
	err = os.repo.CreatOrder(&order)
	if err != nil {
		return err
	}

	for _, item := range cartItems {
		var orderItems models.OrderItem
		orderItems.OrderID = order.ID

		var tempProduct models.Product
		err = os.repo.FindProduct(&tempProduct, item.ProductID)
		if err != nil {
			return err
		}

		orderItems.Price = tempProduct.Price
		orderItems.Quantity = item.Quantity
		orderItems.OrderStatus = "Pending"
		orderItems.ProductID = item.ProductID
		err = os.repo.CreateOrderItem(&orderItems)
		if err != nil {
			return err
		}
		err = os.ProductController(item.ProductID, item.Quantity, tempProduct.Price)
		if err != nil {

			fmt.Println(err)
			os.repo.DeleteOrderItem(order.ID)
			os.repo.DeleteOrder(order.ID)
			return err
		}
	}

	var address models.Address
	address.UserID = uint(userID)
	address.CreatedAt = time.Now()
	address.City = tempOrder.City
	address.Country = tempOrder.Country
	address.OrderID = order.ID
	address.State = tempOrder.State
	address.ZipCode = tempOrder.ZipCode
	address.Street = tempOrder.Street
	err = os.repo.CreatOrderAdress(&address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if tempOrder.Mtherd == "" {
		tempOrder.Mtherd = "cashonpay"
	}

	err = os.repo.DeleteCartItemsByCartID(cart.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
