package app

import (
	"my-gin-app/config"
	"my-gin-app/internal/router"
)

type App interface {
	Start()
}

type impl struct {
	r  router.Router
	db *config.Config
}

func (i *impl) Start() {
	i.r.Start()
}

func NewApp(rout router.Router, db *config.Config) App {
	return &impl{
		r:  rout,
		db: db,
	}
}
