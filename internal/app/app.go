package app

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ndmik-dev/photo-shoot-planner/internal/config"
	"github.com/ndmik-dev/photo-shoot-planner/internal/platform/dbgen"
	"github.com/ndmik-dev/photo-shoot-planner/internal/shoot"
	httptransport "github.com/ndmik-dev/photo-shoot-planner/internal/transport/http"
)

type App struct {
	server *http.Server
}

func New(cfg config.Config, pool *pgxpool.Pool) *App {
	queries := dbgen.New(pool)

	shootRepo := shoot.NewRepository(queries)
	shootService := shoot.NewService(shootRepo)
	shootHandler := shoot.NewHandler(shootService)

	router := httptransport.NewRouter(shootHandler)

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           router,
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
