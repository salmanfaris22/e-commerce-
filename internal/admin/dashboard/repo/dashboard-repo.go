package AdmindashboardRepo

import (
	"my-gin-app/config"
	AdmindashboardInterface "my-gin-app/internal/admin/dashboard/interface"
	"my-gin-app/internal/models"
)

type adminDashBoardrRepoImpl struct {
	config *config.Config
}

func NewAdminDhasBoardpo(config *config.Config) AdmindashboardInterface.AdminDashBoard {
	return &adminDashBoardrRepoImpl{config: config}
}

func (a adminDashBoardrRepoImpl) CountUsers() (int64, error) {
	var totalUser int64
	err := a.config.DB.Model(&models.User{}).Count(&totalUser).Error
	if err != nil {
		return 0, err
	}
	return totalUser, nil
}

func (a adminDashBoardrRepoImpl) CountProducts() (int64, error) {
	var totalProduct int64
	err := a.config.DB.Model(&models.Product{}).Count(&totalProduct).Error
	if err != nil {
		return 0, err
	}
	return totalProduct, nil
}
func (a adminDashBoardrRepoImpl) CountOrders() (int64, error) {
	var totalOrder int64
	err := a.config.DB.Model(&models.Order{}).Count(&totalOrder).Error
	if err != nil {
		return 0, err
	}
	return totalOrder, nil
}
func (a adminDashBoardrRepoImpl) CountTotalProductsSold() (int64, error) {
	var totalOrderSale int64
	err := a.config.DB.Model(&models.OrderItem{}).Count(&totalOrderSale).Error
	if err != nil {
		return 0, err
	}
	return totalOrderSale, nil
}
func (a adminDashBoardrRepoImpl) GetOrderStatusCounts() ([]models.StatusCount, error) {
	var statusCounts []models.StatusCount
	err := a.config.DB.Model(&models.OrderItem{}).Select("order_status, COUNT(*) AS count").Group("order_status").Scan(&statusCounts).Error

	if err != nil {
		return statusCounts, err
	}
	return statusCounts, nil
}
func (a adminDashBoardrRepoImpl) CalculateTotalProfit() (float64, error) {
	var totalSum float64

	if err := a.config.DB.Model(&models.Order{}).Select("SUM(total_Price)").Where("status != ?", "canceled").Scan(&totalSum).Error; err != nil {

		return totalSum, err
	}
	return totalSum, nil
}
func (a adminDashBoardrRepoImpl) ProductSummers(productSummaries *[]models.ProductSummary) error {
	return a.config.DB.Model(&models.OrderItem{}).
		Select("product_id, SUM(quantity) AS total_quantity, SUM(price) AS total_price, SUM(price * quantity) AS total").
		Group("product_id").Scan(&productSummaries).Error
}
func (a adminDashBoardrRepoImpl) FindProduct(tempProduct *models.Product, id uint) error {
	return a.config.DB.Where("id=?", id).First(&tempProduct).Error
}

func (a adminDashBoardrRepoImpl) GetProductSalesByBrand() ([]models.ProductSales, error) {
	var sales []models.ProductSales

	err := a.config.DB.Table("order_items").
		Select("products.brand, SUM(order_items.quantity) AS total_sold, SUM(order_items.quantity * order_items.price) AS total_revenue").
		Joins("JOIN products ON products.id = order_items.product_id").
		Group("products.brand").
		Scan(&sales).Error

	return sales, err
}
