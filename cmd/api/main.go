package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/okok-student-manager/internal/bootstrap"
	"github.com/okok-student-manager/internal/config"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	app, err := bootstrap.New(cfg)
	if err != nil {
		log.Fatalf("initialize application: %v", err)
	}
	defer func() { _ = app.Logger.Sync() }()

	app.Logger.Info("starting HTTP server", zap.String("address", cfg.Address()))
	if err := app.Server.Start(cfg.Address()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.Logger.Fatal("HTTP server stopped", zap.Error(err))
	}
}
