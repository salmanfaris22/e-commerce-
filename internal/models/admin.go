package models

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"unique;type:varchar(100)" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Phone    string `json:"phone" validate:"required,min=10"`
}
