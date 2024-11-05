package adminproduct

import (
	"my-gin-app/internal/models"
	"my-gin-app/pkg/validation"
)

type adminProductServiesImpl struct {
	repo AdminProductRepo
}

func NewAdminProductServeces(repo AdminProductRepo) AdminProductServices {
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
	return nil
}
func (aps adminProductServiesImpl) UpdateProduct(updates interface{}, id string) error {
	err := aps.repo.UpdateProdutcs(updates, id)
	if err != nil {
		return err
	}
	return nil
}
func (aps adminProductServiesImpl) DeleteProduct(id string) error {
	err := aps.repo.DeleteProductRepo(id)
	if err != nil {
		return err
	}
	return nil
}
