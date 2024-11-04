package adminproduct

import "my-gin-app/config"

type adminProductImpl struct {
	config *config.Config
}

func NewAdminProductReposetries(config *config.Config) AdminProductRepo {
	return &adminProductImpl{config: config}
}
