package AdmindashboardServices

import (
	AdmindashboardInterface "my-gin-app/internal/admin/dashboard/interface"
	"my-gin-app/internal/models"
)

type admindashBoardServicesoImpl struct {
	repo AdmindashboardInterface.AdminDashBoard
}

func NewAdminOrdeServices(repo AdmindashboardInterface.AdminDashBoard) AdmindashboardInterface.AdminDashBoardServies {
	return &admindashBoardServicesoImpl{repo: repo}
}

func (ads admindashBoardServicesoImpl) AdminDashBoardServices() (models.DashboardResponse, error) {
	var response models.DashboardResponse

	CountUsers, err := ads.repo.CountUsers()
	if err != nil {
		return response, err
	}
	response.TotalUsers = CountUsers
	CountProducts, err := ads.repo.CountProducts()
	if err != nil {
		return response, err
	}

	response.TotalProducts = CountProducts
	CountOrders, err := ads.repo.CountOrders()
	if err != nil {
		return response, err
	}
	response.TotalOrders = CountOrders
	CountTotalProductsSold, err := ads.repo.CountTotalProductsSold()
	if err != nil {
		return response, err
	}
	response.TotalProductSold = CountTotalProductsSold
	GetOrderStatusCounts, err := ads.repo.GetOrderStatusCounts()
	if err != nil {
		return response, err
	}

	response.OrderStatus = GetOrderStatusCounts
	CalculateTotalProfit, err := ads.repo.CalculateTotalProfit()
	if err != nil {
		return response, err
	}

	response.TotalProfit = CalculateTotalProfit
	ProductAnalist, err := ads.ProductAnalysis()
	if err != nil {
		return response, err
	}
	sales, err := ads.repo.GetProductSalesByBrand()
	if err != nil {
		return response, err
	}
	response.ProductSales = sales
	response.ProductAnalysis = ProductAnalist
	return response, nil
}

func (ads admindashBoardServicesoImpl) ProductAnalysis() ([]models.ProductAnalist, error) {
	var productSummaries []models.ProductSummary
	var Getalldetails []models.ProductAnalist

	err := ads.repo.ProductSummers(&productSummaries)
	if err != nil {
		return Getalldetails, err
	}
	for _, item := range productSummaries {
		var tempModel models.ProductAnalist
		tempModel.ProductID = item.ProductID
		tempModel.TotalPrice = item.Total
		tempModel.TotalQuantity = item.TotalQuantity
		tempModel.Total = item.Total

		var tempProduct models.Product
		err = ads.repo.FindProduct(&tempProduct, item.ProductID)
		if err != nil {

			return Getalldetails, err
		}

		tempModel.Brand = tempProduct.Brand
		tempModel.CompanyName = tempProduct.CompanyName
		tempModel.Name = tempProduct.Name
		Getalldetails = append(Getalldetails, tempModel)
	}

	return Getalldetails, nil
}
