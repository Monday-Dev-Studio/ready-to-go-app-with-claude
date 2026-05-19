package main

import (
	"log/slog"
	"os"

	apphttp "app/internal/delivery/http"
	"app/pkg/config"
	"app/pkg/database"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "err", err)
		os.Exit(1)
	}

	db, err := database.NewPool(cfg.DB.DSN())
	if err != nil {
		slog.Error("connect to database", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	app := apphttp.NewApp(cfg, db)

	slog.Info("server starting", "port", cfg.App.Port)
	if err = app.Listen(":" + cfg.App.Port); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
