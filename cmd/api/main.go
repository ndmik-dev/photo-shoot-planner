package main

import (
	"context"
	"log"
	"time"

	"github.com/ndmik-dev/photo-shoot-planner/internal/app"
	"github.com/ndmik-dev/photo-shoot-planner/internal/config"
	"github.com/ndmik-dev/photo-shoot-planner/internal/platform/db"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	log.Println("database connected")

	application := app.New(cfg, pool)

	log.Printf("starting server on :%s", cfg.AppPort)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
