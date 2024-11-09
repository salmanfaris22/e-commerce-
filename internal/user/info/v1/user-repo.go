package userInfo

import "my-gin-app/config"

type UserRepo interface {
}

type userInfoImpl struct {
	config *config.Config
}

func NewUserRepo(config *config.Config) UserRepo {
	return &userInfoImpl{config: config}
}
