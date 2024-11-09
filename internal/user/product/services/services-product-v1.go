package productServices

import (
	"errors"
	"strconv"

	"my-gin-app/internal/models"
	productInterface "my-gin-app/internal/user/product/interface"
)

type producctServicesImpl struct {
	Productrepo productInterface.ProductRepo
}

func NewProducctServicesV1(repo productInterface.ProductRepo) productInterface.ProducctServices {
	return &producctServicesImpl{Productrepo: repo}
}

func (ps producctServicesImpl) GetAllProduct() ([]models.Product, error) {
	var product []models.Product
	err := ps.Productrepo.GetAllProductModel(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (ps producctServicesImpl) GetIDProductService(productID string) (models.Product, error) {
	var product models.Product
	id, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		return product, err
	}
	product.ID = uint(id)
	err = ps.Productrepo.GetProductModelById(&product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (ps producctServicesImpl) SerchProductService(searchItem string) ([]models.Product, error) {
	var products []models.Product
	err := ps.Productrepo.SearchProductRepo(&products, searchItem)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (ps producctServicesImpl) FilterProduct(filter models.Filter, Available, maxPriceStr, minPriceStr string) ([]models.Product, error) {
	var products []models.Product
	if Available == "true" {
		*filter.IsAvailable = true
	} else if Available == "false" {
		*filter.IsAvailable = false
	}
	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			return products, errors.New("valid min price")
		}
		filter.MinPrice = &minPrice
	}
	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			return products, errors.New("valid min price")
		}
		filter.MaxPrice = &maxPrice
	}
	query := ps.Productrepo.FilterQuery()
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if Available != "" {
		query = query.Where("is_available = ?", *filter.IsAvailable)
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.Brand != "" {
		query = query.Where("brand = ?", filter.Brand)
	}

	err := ps.Productrepo.QueryFindProduct(query, &products)
	if err != nil {
		return products, err
	}
	return products, nil
}
