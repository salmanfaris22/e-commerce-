package models

type DashboardResponse struct {
	TotalUsers       int64            `json:"total_users"`
	TotalProducts    int64            `json:"total_products"`
	TotalOrders      int64            `json:"total_orders"`
	TotalProductSold int64            `json:"total_product_sold"`
	OrderStatus      []StatusCount    `json:"order_status"`
	TotalProfit      float64          `json:"total_profit"`
	ProductAnalysis  []ProductAnalist `json:"product_analysis"`
	ProductSales     []ProductSales
}
type ProductSales struct {
	Brand        string  `json:"brand"`
	TotalSold    int     `json:"total_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}
