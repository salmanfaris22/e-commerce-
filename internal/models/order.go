package models

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_Price" gorm:"not null" `
	Status     string      `json:"status" gorm:"default:'pending'"`
	CreatedAt  time.Time   `json:"created_at" `
	UpdatedAt  time.Time   `json:"updated_at"`
	Addresses  []Address   `gorm:"foreignKey:OrderID" json:"addresses" validate:"dive"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"order_items" validate:"dive"`
}

type OrderItem struct {
	ID          uint    `gorm:"primaryKey" json:"id" validate:"required"`
	OrderID     uint    `json:"order_id" validate:"required"`
	ProductID   uint    `json:"product_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"gte=1"`
	Price       float64 `json:"price" validate:"gte=0"`
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	OrderStatus string  `json:"order_status" gorm:"default:'pending'"`
}

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `json:"order_id" validate:"required"`
	Amount        float64   `json:"amount" validate:"gte=0"`
	Method        string    `json:"method" gorm:"not null" validate:"required"` // e.g., "Credit Card", "PayPal", "Bank Transfer"
	Status        string    `json:"status" gorm:"default:'pending'"`            // e.g., "pending", "completed", "failed"
	TransactionID string    `json:"transaction_id"`                             // Optional for transaction reference
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DemoOrder struct {
	UserID        uint           `json:"user_id" validate:"required"`
	Street        string         `json:"street" gorm:"not null" validate:"required"`
	City          string         `json:"city" gorm:"not null" validate:"required"`
	State         string         `json:"state" gorm:"not null" validate:"required"`
	Pymentmetherd string         `json:"pyment_metherd"validate:"required"`
	ZipCode       string         `json:"zip_code" gorm:"not null" validate:"required"`
	Country       string         `json:"country" gorm:"not null" validate:"required"`
	OrderItem     []OrderItem    `json:"order_items" validate:"dive,required"`
	ProductName   string         `json:"product_name" `
	Images        pq.StringArray `json:"images"`
}

type StatusCount struct {
	OrderStatus string `json:"order_status"`
	Count       int    `json:"count"`
}

type ProductSummary struct {
	ProductID     uint    `json:"product_id"`
	TotalQuantity int64   `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
	Total         float64 `json:"total"`
}

type ProductAnalist struct {
	ProductID     uint    `json:"product_id"`
	TotalQuantity int64   `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
	Total         float64 `json:"total"`
	Name          string  `json:"name"`
	CompanyName   string  `json:"company_name"`
	Brand         string  `json:"brand"`
}
