package bot

import (
	"github.com/EgorKo25/common/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DEBUG = iota
	PRODUCTION
)

type Bot struct {
	API          *tgbotapi.BotAPI
	UpdateConfig tgbotapi.UpdateConfig
	logger       logger.ILogger
}

func New(token string, pollingTO int, mode int, l *logger.Logger) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	switch mode {
	case DEBUG:
		botAPI.Debug = true
	case PRODUCTION:
		botAPI.Debug = false
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = pollingTO

	return &Bot{API: botAPI, UpdateConfig: u, logger: l}, nil
}
