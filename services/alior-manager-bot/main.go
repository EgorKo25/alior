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
		l.Fatal(err.Error())
	}

	// Инициализация брокера
	err = broker.NewBroker(cfg.Broker)
	if err != nil {
		l.Fatal(err.Error())
	}

	// Инициализация бота
	tgBot, err := bot.New(
		cfg.Bot.BotToken,
		cfg.Bot.BotPolingTO,
		bot.DEBUG,
		log,
		bot.WithPublisherName("ask_publisher"),
		bot.WithConsumerName("ans_consumer"),
	)
	if err != nil {
		log.Fatal("Failed to initialize bot: ", err)
	}

	// Запуск бота
	if err := tgBot.Run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
