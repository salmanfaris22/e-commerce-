package order

import "my-gin-app/config"

type orderRepoImpl struct {
	config config.Config
}

func NewOrderRepoV1(config config.Config) OrderRepo {
	return &orderRepoImpl{config: config}
}
