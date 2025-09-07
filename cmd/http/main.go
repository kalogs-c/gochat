package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kalogs-c/gochat/config"
	"github.com/kalogs-c/gochat/internal/httpserver"
	"github.com/kalogs-c/gochat/internal/storage/sqlite"
	sqlc "github.com/kalogs-c/gochat/sql/sqlc_generated"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := sqlite.MustConnect(ctx, "./gochat.db")
	queries := sqlc.New(db)

	server := httpserver.NewHTTPServer(
		logger,
		queries,
		&config.Config{Host: "localhost", Port: "42069"},
		httpserver.WithLogging(logger),
	)

	go server.MustServe()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	logger.Info("Server gracefully stopped")
}
