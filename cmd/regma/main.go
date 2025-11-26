package main

import (
	"log/slog"
	"os"

	"Regma/internal/api"
	"Regma/internal/config"
)

func main() {
	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Config
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// TODO: Initialize database
	// db, err := storage.NewPostgres(cfg.DatabaseURL)

	// TODO: Initialize repositories
	// investorRepo := storage.NewInvestorRepository(db)

	// TODO: Initialize services
	// investorSvc := investor.NewService(investorRepo)

	// Server
	server := api.NewServer(cfg, logger)

	slog.Info("starting server", "port", cfg.Port)
	if err := server.Run(); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}
