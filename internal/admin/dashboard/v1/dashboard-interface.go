package Admindashboard

import (
	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type AdminDashBoard interface {
	CountUsers() (int64, error)
	CountProducts() (int64, error)
	CountOrders() (int64, error)
	CountTotalProductsSold() (int64, error)
	GetOrderStatusCounts() ([]models.StatusCount, error)
	CalculateTotalProfit() (float64, error)
	ProductSummers(productSummaries *[]models.ProductSummary) error
	FindProduct(tempProduct *models.Product, id uint) error
	GetProductSalesByBrand() ([]models.ProductSales, error)
}
type AdminDashBoardInterface interface {
	AdmindashBoardGetAll(ctx *gin.Context)
}
type AdminDashBoardServies interface {
	AdminDashBoardServices() (models.DashboardResponse, error)
}
