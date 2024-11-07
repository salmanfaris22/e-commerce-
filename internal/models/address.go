package models

import "time"

type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id" validate:"required"`
	UserID    uint      `json:"user_id" validate:"required"`
	OrderID   uint      `json:"order_id" validate:"required"`
	Street    string    `json:"street" gorm:"not null" validate:"required"`
	City      string    `json:"city" gorm:"not null" validate:"required"`
	State     string    `json:"state" gorm:"not null" validate:"required"`
	ZipCode   string    `json:"zip_code" gorm:"not null" validate:"required"`
	Country   string    `json:"country" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
type GetOrderdetils struct {
	ID        uint      `gorm:"primaryKey" json:"id" validate:"required"`
	UserID    uint      `json:"user_id" validate:"required"`
	OrderID   uint      `json:"order_id" validate:"required"`
	Street    string    `json:"street" gorm:"not null" validate:"required"`
	City      string    `json:"city" gorm:"not null" validate:"required"`
	State     string    `json:"state" gorm:"not null" validate:"required"`
	ZipCode   string    `json:"zip_code" gorm:"not null" validate:"required"`
	Country   string    `json:"country" gorm:"not null" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	Mtherd    string    `json:"methord" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
