package bot

import (
	"context"
	"github.com/EgorKo25/common/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	DEBUG = iota
	PRODUCTION
)

type HandlerFunc func(ctx context.Context, update *tgbotapi.Update) error

type Bot struct {
	API           *tgbotapi.BotAPI
	CommandConfig tgbotapi.SetMyCommandsConfig
	UpdateConfig  tgbotapi.UpdateConfig
	logger        logger.ILogger
	handlers      map[string]HandlerFunc
	PublisherName string
	ConsumerName  string
}

func New(token string, pollingTO int, mode int, l *logger.Logger, publisherName, consumerName string) (*Bot, error) {
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

	bot := &Bot{
		API:           botAPI,
		UpdateConfig:  tgbotapi.NewUpdate(0),
		logger:        l,
		handlers:      make(map[string]HandlerFunc),
		PublisherName: publisherName,
		ConsumerName:  consumerName,
	}

	bot.UpdateConfig.Timeout = pollingTO

	bot.CommandConfig, _ = setupCommands()

	if _, err := bot.API.Request(bot.CommandConfig); err != nil {
		return nil, err
	}

	bot.initHandlers()

	return bot, nil
}
