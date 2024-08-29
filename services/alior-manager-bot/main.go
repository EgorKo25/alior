package main

import (
	"alior-manager-bot/src/config"
	"alior-manager-bot/src/core"
	"github.com/EgorKo25/common/logger"

	l "log"
)

func main() {
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

	// Инициализация бота
	b, err := bot.New(cfg.Bot.BotToken, cfg.Bot.BotPolingTO, bot.DEBUG, log)
	if err != nil {
		log.Fatal("Failed to initialize bot")
	}

	commands := b.CommandConfig.Commands
	for _, command := range commands {
		l.Printf("command: %s", command)
	}

	// Запуск бота
	if err := b.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
