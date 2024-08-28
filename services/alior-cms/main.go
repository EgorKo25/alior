package main

import (
	"callback_service/src/broker"
	"callback_service/src/config"
	"callback_service/src/database"
	"callback_service/src/service"
	"context"
	"github.com/EgorKo25/common/logger"
	l "log"
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
	log, err := logger.NewLogger(logger.PRODUCTION)
	if err != nil {
		l.Fatal(err)
	}

	// Инициализация БД
	db, err := database.New(ctx, cfg)
	if err != nil {
		log.Fatal("failed to initialize database: %v", err)
	}
	defer db.Close()

	// Инициализация брокера
	b, err := broker.NewBroker(cfg.MsgBroker.URL, log)
	if err != nil {
		log.Fatal("failed to initialize broker: %v", err)
	}

	// Инициализация сервиса
	cms := service.NewCMS(b, db, log)

	// Запуск приложения
	if err := cms.Run(ctx); err != nil {
		log.Fatal("failed to run CMS service: %v", err)
	}
}
