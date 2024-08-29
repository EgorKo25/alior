package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Run() error {
	errCh := make(chan error, 1)

	go func() {
		updates := b.API.GetUpdatesChan(b.UpdateConfig)

		for update := range updates {
			if update.Message == nil {
				continue
			}

			if !update.Message.IsCommand() {
				continue
			}

			b.logger.Info("Got message: %s", update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}

			if _, err := b.API.Send(msg); err != nil {
				b.logger.Fatal("Failed to send message: %s", err)
			} else {
				b.logger.Info("Message sent")
			}
		}
	}()

	return <-errCh
}
