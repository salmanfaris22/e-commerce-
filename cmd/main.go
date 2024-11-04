package main

import (
	"log"
	"my-gin-app/config"
	"my-gin-app/internal/app"
	"my-gin-app/internal/router"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Could not initialize application config: %v", err)
	}
	r := router.NewRouter(cfg)
	ap := app.NewApp(r, cfg)
	ap.Start()
}
