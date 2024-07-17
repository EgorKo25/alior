package main

import (
	"callback_service/src/config"
	"callback_service/src/migrator"
	"callback_service/src/repository"
	"callback_service/src/service"
	"callback_service/src/transport/amqp"
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
	// Инициализация context
	ctx := context.Background()

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

	// Иициализация БД
	pool, err := pgxpool.New(ctx, cfg.Database.Url)
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
	err = amqp.Consume(ctx, cfg.MsgBroker.Url, "create", svc)
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
