package models

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	ProductID uint      `json:"product_id"`
	Rating    int       `json:"rating" gorm:"not null"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
