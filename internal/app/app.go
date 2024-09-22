package app

import (
	"context"

	grpcapp "github.com/DaniilZ77/authorization/internal/app/gprc"
	"github.com/DaniilZ77/authorization/internal/config"
	"github.com/DaniilZ77/authorization/internal/lib/logger"
	"github.com/DaniilZ77/authorization/internal/lib/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(cfg *config.Config) *App {
	ctx := context.Background()

	// Init logger
	logger.New(cfg.Log.Level)

	// Postgres connection
	pg, err := postgres.New(ctx, cfg.DB.URL)
	if err != nil {
		logger.Log().Fatal(ctx, "error with connection to database: %s", err.Error())
	}
	defer pg.Close(ctx)

	// Service
	//

	gRPCApp := grpcapp.New(cfg.Port)

	return &App{
		GRPCServer: gRPCApp,
	}
}
