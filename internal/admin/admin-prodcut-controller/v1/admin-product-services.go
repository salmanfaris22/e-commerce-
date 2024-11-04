package adminproduct

type adminProductServiesImpl struct {
	repo AdminProductRepo
}

func NewAdminProductServeces(repo AdminProductRepo) AdminProductServices {
	return &adminProductServiesImpl{repo: repo}
}
