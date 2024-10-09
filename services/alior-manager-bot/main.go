package main

import (
	"alior-manager-bot/src/config"
	"alior-manager-bot/src/core"
	"context"
	"github.com/EgorKo25/common/broker"
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

	// Инициализация брокера
	err = broker.InitBroker(cfg.Broker.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	// паблишер для отправки запросов
	pubConfig := broker.NewPublisherConfig(cfg.Broker.Exchange.Name, cfg.Broker.Exchange.Kind, "ask")
	if pubConfig == nil {
		log.Fatal("publisher config not created")
	}

	err = broker.CreatePublisher("ask_publisher", pubConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// читатель ответов
	consConfig := broker.NewConsumerConfig(cfg.Broker.Exchange.Name, "ans", "ans", "")
	if consConfig == nil {
		log.Fatal("Consumer config not created")
	}

	err = broker.CreateConsumer("ans_consumer", consConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Инициализация бота
	tgBot, err := bot.New(cfg.Bot.BotToken, cfg.Bot.BotPolingTO, bot.DEBUG, log, "ask_publisher", "ans_consumer")
	if err != nil {
		log.Fatal("Failed to initialize bot")
	}

	// Запуск бота
	if err := tgBot.Run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
