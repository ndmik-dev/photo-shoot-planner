package main

import (
	"log"

	"github.com/ndmik-dev/photo-shoot-planner/internal/app"
	"github.com/ndmik-dev/photo-shoot-planner/internal/config"
)

func main() {
	cfg := config.Load()
	application := app.New(cfg)

	log.Printf("starting server on :%s", cfg.AppPort)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
