package main

import (
	"callback_service/src/config"
	"callback_service/src/logger"
	"callback_service/src/migrator"
	"callback_service/src/repository"
	"callback_service/src/service"
	"callback_service/src/transport/amqp"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// Инициализация context
	ctx := context.Background()

	// Инициализация config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Инициализация logger
	log := SetupLogger(cfg.Env)

	log.Info("Starting Callback service")
	log.Info("env: ", cfg.Env)
	log.Info("Database url", cfg.Database.Url)
	log.Info("MsgBroker url", cfg.MsgBroker.Url)
	log.Info("cfg", cfg)

	// Иициализация БД
	pool, err := pgxpool.New(ctx, cfg.Database.Url)
	if err != nil {
		log.Error("failed to connect to database", err)
	}

	err = database.MigrateDatabase(pool)
	if err != nil {
		log.Error("migrations failed", err)
	}

	// Запуск Consumer'a
	repo := repository.NewRepository(pool)
	svc := service.NewCallbackService(repo)
	err = amqp.Consume(ctx, cfg.MsgBroker.Url, "create", svc)
	if err != nil {
		log.Error("failed to start consumer", err)
	}
}

func SetupLogger(env string) logger.ILogger {
	var log *zap.Logger
	var err error

	switch env {
	case envLocal:
		log, err = logger.NewDevLogger()
		if err != nil {
			panic(err)
		}
	case envProd:
		log, err = logger.NewProdLogger()
		if err != nil {
			panic(err)
		}
	}
	return logger.NewZapLogger(log)
}
