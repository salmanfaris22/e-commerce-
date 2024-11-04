package wishlist

import (
	"strconv"

	"gorm.io/gorm"

	"my-gin-app/internal/models"
)

type wishLisSerivestImpel struct {
	repo WishListRepo
}

func NewWishlistServises(repo WishListRepo) WishListServices {
	return &wishLisSerivestImpel{repo: repo}
}

func (ws wishLisSerivestImpel) GetAllWihslistItems(userId string) (string, []models.Product, error) {
	var products []models.Product
	var wishlists models.Wishlist
	err := ws.repo.GetWishlistItemsAll(userId, &wishlists)
	if err != nil {
		return "you hanve no wishlist", products, err
	}
	for _, item := range wishlists.Items {
		var product models.Product
		newId := strconv.FormatUint(uint64(item.ID), 10)
		err = ws.repo.GetPpduct(&product, newId)
		if err != nil {
			return "no Product avalible", products, err
		}
		products = append(products, product)
	}
	return "", products, nil
}

func (ws wishLisSerivestImpel) WishListAddremove(productID, userID string) (string, error) {
	var product models.Product
	err := ws.repo.GetPpduct(&product, productID)
	if err != nil {
		return "product Can't find product", err
	}
	var wishlist models.Wishlist
	err = ws.repo.FindWishlist(userID, &wishlist)
	if err != nil {
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				uid, _ := strconv.Atoi(userID)

				wishlist.UserID = uint(uid)
				err = ws.repo.CreatWishlist(&wishlist)
				if err != nil {
					return "Can't craete wishlist", err
				}
			} else {
				return "product Can't find wishlist", err
			}

		}

	}
	var listItem models.WishlistItem
	err = ws.repo.FindWishlistItem(wishlist.ID, product.ID, &listItem)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			listItem.ProductID = product.ID
			listItem.WishlistID = wishlist.ID
			err = ws.repo.CreatWishlistItem(&listItem)
			if err != nil {
				return "Can't craete wishlistItem", err
			}
			return "wish list added successfully", nil
		} else {
			return "Can't find list items", err
		}

	}
	ws.repo.DeleteWishlistItem(&listItem)
	if err != nil {
		return "Unable to remove item from wishlist", err
	}
	return "wishlist item removed successfully", nil
}
