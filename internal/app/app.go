package app

import (
	"context"

	grpcapp "github.com/DaniilZ77/authorization/internal/app/gprc"
	"github.com/DaniilZ77/authorization/internal/config"
	"github.com/DaniilZ77/authorization/internal/lib/logger"
	"github.com/DaniilZ77/authorization/internal/lib/postgres"
	"github.com/DaniilZ77/authorization/internal/service/auth"
	"github.com/DaniilZ77/authorization/internal/store/postgres/user"
)

type App struct {
	GRPCServer *grpcapp.App
	PG         *postgres.Postgres
}

func New(ctx context.Context, cfg *config.Config) *App {
	// Init logger
	logger.New(cfg.Log.Level)

	// Postgres connection
	pg, err := postgres.New(ctx, cfg.DB.URL)
	if err != nil {
		logger.Log().Fatal(ctx, "error with connection to database: %s", err.Error())
	}

	// Store
	userStore := user.New(pg)

	// Service
	authService := auth.New(userStore)

	// gRPC server
	gRPCApp := grpcapp.New(ctx, authService, cfg)

	return &App{
		GRPCServer: gRPCApp,
		PG:         pg,
	}
}
