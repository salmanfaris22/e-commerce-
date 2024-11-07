package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my-gin-app/internal/models"
)

// connectDB initializes a database connection using GORM and returns a pointer to the DB instance.
func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=poomon dbname=newEcommers port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Perform automatic migration for all the models.
	err = db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.Product{},
		&models.Review{},
		&models.Wishlist{},
		&models.WishlistItem{},
		&models.OrderItem{},
		&models.Admin{},
		&models.UserToken{},
		&models.ProductImage{},
		&models.Payment{},
	)
	if err != nil {
		return nil, fmt.Errorf("can't AutoMigrate: %w", err)
	}

	fmt.Println("Database connected and migrated successfully.")
	return db, nil
}

type Config struct {
	DB *gorm.DB
}

func NewConfig() (*Config, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: db,
	}, nil
}

// Uncomment this function if you want to provide a way to retrieve the DB outside of Config.
// func GetDB() *gorm.DB {
// 	return db
// }
