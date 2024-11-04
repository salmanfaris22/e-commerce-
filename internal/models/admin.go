package models

type Admin struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string `gorm:"unique;type:varchar(100)" json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Password  string `json:"password" validate:"required,min=8,max=255"`
	Phone     string `json:"phone" validate:"required,min=10"`
}
