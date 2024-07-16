package main

import (
	database "callback_service/cmd/migrator"
	"callback_service/internal/config"
	"callback_service/internal/repository"
	"callback_service/internal/service"
	"callback_service/internal/transport/amqp"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// Инициализация config
	cfg := config.MustLoad()

	// Инициализация logger
	log := setupLogger(cfg.Env)
	log.Info("Starting Callback service",
		slog.String("env", cfg.Env),
		slog.String("Database url", cfg.Database.Url),
		slog.String("MsgBroker url", cfg.MsgBroker.Url),
		slog.Any("cfg", cfg),
	)
	log.Debug("Debug msg")
	log.Error("Error msg")
	log.Warn("Warning msg")

	// Иициализация БД
	pool, err := pgxpool.New(context.Background(), cfg.Database.Url)
	if err != nil {
		log.Error("Failed to connect to database: %v", err)
	}

	err = database.MigrateDatabase(pool)
	if err != nil {
		log.Error("Migrations failed: %v", err)
	}

	// Запуск Consumer'a
	repo := repository.NewRepository(pool)
	svc := service.NewCallbackService(repo)
	err = amqp.Consume(cfg.MsgBroker.Url, "create", svc)
	if err != nil {
		log.Error("Failed to start consumer: %v", err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
