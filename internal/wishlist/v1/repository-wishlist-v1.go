package wishlist

import (
	"my-gin-app/config"
	"my-gin-app/internal/models"
)

type wishListImpel struct {
	config config.Config
}

func NewWishListrepo(config *config.Config) WishListRepo {
	return &wishListImpel{config: *config}
}

func (wr wishListImpel) GetPpduct(product *models.Product, productId string) error {
	return wr.config.DB.Preload("Images").First(&product, productId).Error
}

func (wr wishListImpel) FindWishlist(userID string, wishlist *models.Wishlist) error {
	return wr.config.DB.Where("user_id=?", userID).First(&wishlist).Error
}
func (wr wishListImpel) CreatWishlist(wishlist *models.Wishlist) error {
	return wr.config.DB.Create(&wishlist).Error
}
func (wr wishListImpel) FindWishlistItem(wishlistId uint, productId uint, listItem *models.WishlistItem) error {
	return wr.config.DB.Where("wishlist_id=? AND product_id=?", wishlistId, productId).First(&listItem).Error
}
func (wr wishListImpel) CreatWishlistItem(listItem *models.WishlistItem) error {
	return wr.config.DB.Create(&listItem).Error
}

func (wr wishListImpel) DeleteWishlistItem(listItem *models.WishlistItem) error {
	return wr.config.DB.Delete(&listItem).Error
}

func (wr wishListImpel) GetWishlistItemsAll(userID string, wishlists *models.Wishlist) error {
	return wr.config.DB.Preload("Items").Where("user_id = ?", userID).First(&wishlists).Error
}
func (wr wishListImpel) GetWishlistItemsAllitem(id uint, wishlists *[]models.WishlistItem) error {
	return wr.config.DB.Where("wishlist_id = ?", id).Find(&wishlists).Error
}
func (wr wishListImpel) FindProductById(product *models.Product, productId uint) error {
	return wr.config.DB.Preload("Images").Where("id=?", productId).First(&product).Error
}
