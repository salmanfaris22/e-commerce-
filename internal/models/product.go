package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// UserID      uint           `json:"user_id"`
	Name        string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Description string         `json:"description" gorm:"not null" validate:"required,min=10,max=500"`
	Price       float64        `json:"price" gorm:"not null" validate:"required,min=0"`
	Stock       int            `json:"stock" gorm:"not null" validate:"required,min=0"`
	IsAvailable bool           `json:"is_available" gorm:"default:true"`
	CompanyName string         `json:"company_name" validate:"omitempty,min=2,max=100"`
	Brand       string         `json:"brand" validate:"omitempty,min=2,max=50"`
	Size        pq.StringArray `gorm:"type:text[]" json:"size" validate:"dive,required"`
	Category    string         `json:"category" validate:"omitempty,min=2,max=50"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Images      []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
}

type ProductImage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `json:"product_id" gorm:"not null" validate:"required"`
	URL       string    `json:"url" gorm:"not null" validate:"required,url"`
	IsMain    bool      `json:"is_main" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Filter struct {
	MinPrice    *float64 `form:"min_price"`
	MaxPrice    *float64 `form:"max_price"`
	IsAvailable *bool    `form:"is_available"`
	Category    string   `form:"category"`
	Brand       string   `form:"brand"`
}

func (p Product) FindProduct(db *gorm.DB, ProductID uint64, product *Product) error {
	err := db.First(&product, ProductID).Error
	if err != nil {
		return err
	}
	return nil
}
