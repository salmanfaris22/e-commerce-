package adminproductServices

import (
	"fmt"
	"time"

	adminproductInterFace "my-gin-app/internal/admin/admin-prodcut-controller/interface"
	"my-gin-app/internal/models"
	"my-gin-app/pkg/validation"
)

type adminProductServiesImpl struct {
	repo adminproductInterFace.AdminProductRepo
}

func NewAdminProductServeces(repo adminproductInterFace.AdminProductRepo) adminproductInterFace.AdminProductServices {
	return &adminProductServiesImpl{repo: repo}
}

func (aps adminProductServiesImpl) AddProduct(product models.Product) error {
	err := validation.ValidateUser(product)
	if err != nil {
		return err
	}
	err = aps.repo.AddProduct(&product)
	if err != nil {
		return err
	}
	for _, img := range product.Images {
		img.CreatedAt = time.Now()
		img.ProductID = product.ID

		if img.ID == 0 {
			err = aps.repo.SaveIMg(&img)
			if err != nil {
				return err
			}
		} else {

			err = aps.repo.UpdateImges(&img)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (aps adminProductServiesImpl) UpdateProduct(product models.Product, id string) error {
	var existingProduct models.Product
	err := aps.repo.FindProduct(id, &existingProduct)
	if err != nil {
		return err
	}
	for _, v := range existingProduct.Images {
		err = aps.repo.DeleteImaged(v.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println(product.Images)
	for _, img := range product.Images {
		err = aps.repo.UpdateImges(&img)
		if err != nil {
			return err
		}
	}

	return aps.repo.UpdateProdutcs(&product)
}
func (aps adminProductServiesImpl) DeleteProduct(id string) error {
	var existingProduct models.Product
	err := aps.repo.FindProduct(id, &existingProduct)
	if err != nil {
		return err
	}
	var images []models.ProductImage
	err = aps.repo.FindAllImages(existingProduct.ID, &images)
	if err != nil {
		return err
	}
	for _, v := range images {
		err = aps.repo.DeleteImaged(v.ID)
		if err != nil {
			return err
		}
	}
	return aps.repo.DeleteProductRepo(id)
}
