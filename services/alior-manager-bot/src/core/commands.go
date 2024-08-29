package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupCommands() (tgbotapi.SetMyCommandsConfig, error) {
	commandConfig := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "get_initial_callback",
			Description: "Запрашивает самую старую заявку"},
	)
	return commandConfig, nil
}
