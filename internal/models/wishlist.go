package models

import (
	"time"
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
	WishlistID uint `json:"wishlist_id"`
	ProductID  uint `json:"product_id"`
}
