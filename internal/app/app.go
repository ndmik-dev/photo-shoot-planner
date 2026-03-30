package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ndmik-dev/photo-shoot-planner/internal/config"
	router "github.com/ndmik-dev/photo-shoot-planner/internal/transport/http"
)

type App struct {
	server *http.Server
}

func New(cfg config.Config) *App {
	handler := router.NewRouter()

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &App{server: srv}
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
