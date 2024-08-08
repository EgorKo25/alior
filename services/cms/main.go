package main

import (
	"callback_service/src/broker"
	"callback_service/src/config"
	"callback_service/src/database"
	"callback_service/src/logger"
	"callback_service/src/service"
	"context"
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
	log := logger.NewZapLogger()
	log.Info("Config: ", cfg)

	// Инициализация БД
	db, err := database.New(ctx, cfg)
	if err != nil {
		log.Fatal("failed to initialize database: %v", err)
	}
	defer db.Close()

	// Инициализация брокера
	b := broker.NewBroker(cfg.MsgBroker.Url, log)

	// Инициализация сервиса
	cms := service.NewCMS(b, db, log)

	// Запуск приложения
	if err := cms.Run(ctx); err != nil {
		log.Fatal("failed to run CMS service: %v", err)
	}
}
