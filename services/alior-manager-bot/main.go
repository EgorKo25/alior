package main

import (
	"alior-manager-bot/src/config"
	"github.com/EgorKo25/common/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	l "log"
)

func main() {
	log, err := logger.NewLogger(logger.PRODUCTION)
	if err != nil {
		l.Fatal("Failed to initialize logger")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.BotToken)
	if err != nil {
		log.Fatal("Failed to initialize bot: %s", err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				log.Error("failed to send message: %s", err)
			}
		}
	}
}
