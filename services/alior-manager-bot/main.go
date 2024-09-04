package main

import (
	"alior-manager-bot/src/config"
	"alior-manager-bot/src/core"
	"alior-manager-bot/src/transport/broker"
	"context"
	"github.com/EgorKo25/common/logger"

	l "log"
)

func main() {
	ctx := context.Background()

	// Инициализация логгера
	log, err := logger.NewLogger(logger.PRODUCTION)
	if err != nil {
		l.Fatal("Failed to initialize logger")
	}

	// Загрузка конфига
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Создаем новый экземпляр брокера
	b, err := broker.NewBroker(
		"amqp://guest:guest@localhost:5672/", // URI подключения
		"test-exchange",                      // Имя Exchange
		"direct",                             // Тип Exchange
		"test-key",                           // Routing Key
		"testin-queue",                       // Имя очереди
		log,                                  // Логгер
	)
	if err != nil {
		log.Error("failed to create broker service: %s", err)
	}

	// Инициализация бота
	tgBot, err := bot.New(cfg.Bot.BotToken, cfg.Bot.BotPolingTO, bot.DEBUG, log, b)
	if err != nil {
		log.Fatal("Failed to initialize bot")
	}

	// Запуск бота
	if err := tgBot.Run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
