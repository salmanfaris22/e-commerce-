package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-gin-app/internal/models"
)

type productImpl struct {
	productService ProducctServices
}

func NewProductHandlerV1(services ProducctServices) ProductHandle {
	return &productImpl{productService: services}
}

func (ph productImpl) GetAllProduct(ctx *gin.Context) {
	product, err := ph.productService.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Product", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": product,
	})
}

func (ph productImpl) GetProductById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	product, err := ph.productService.GetIDProductService(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Product", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": product,
	})
}

func (ph productImpl) SerchProduct(ctx *gin.Context) {
	searchItem := ctx.Query("product")
	products, err := ph.productService.SerchProductService(searchItem)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Product", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": products,
	})
}

func (ph productImpl) FilterProducts(ctx *gin.Context) {
	var filter models.Filter

	filter.MinPrice = nil
	filter.MaxPrice = nil
	filter.IsAvailable = new(bool)
	filter.Brand = ctx.Query("brand")
	filter.Category = ctx.Query("category")
	Available := ctx.Query("is_available")
	maxPriceStr := ctx.Query("max_price")
	minPriceStr := ctx.Query("min_price")

	products, err := ph.productService.FilterProduct(filter, Available, maxPriceStr, minPriceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid Product", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": products,
	})
}
