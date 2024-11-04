package models

import (
	"time"

	"github.com/lib/pq"
)

type Wishlist struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `json:"user_id"` //foreignKey  user
	Items     []WishlistItem `gorm:"foreignKey:WishlistID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type WishlistItem struct {
	ID         uint `gorm:"primaryKey"`
	WishlistID uint `json:"wishlist_id"` //foreignKey wish list
	ProductID  uint `json:"product_id"`
}

type ItemsWithDetails struct {
	ItemID      uint           `json:"item_id"`
	ProductID   uint           `json:"product_id"`
	ProductName string         `json:"product_name"`
	Price       float64        `json:"price"`
	Images      pq.StringArray `json:"images"`
}
