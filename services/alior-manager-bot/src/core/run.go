package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		updates := b.API.GetUpdatesChan(b.UpdateConfig)

		for update := range updates {
			if update.Message != nil && update.Message.IsCommand() {
				b.logger.Info("Got command: %s", update.Message.Text)
				command := update.Message.Command()

				if handler, exists := b.handlers[command]; exists {
					if err := handler(ctx, &update); err != nil {
						b.logger.Error("handler error: %v", err)
						errCh <- err
					}
				} else {
					b.unknownCommandHandler(&update)
				}
			}

			if update.CallbackQuery != nil {
				callbackData := update.CallbackQuery.Data
				b.logger.Info("Got callback query: %s", callbackData)

				if handler, exists := b.handlers[callbackData]; exists {
					if err := handler(ctx, &update); err != nil {
						b.logger.Error("callback handler error: %v", err)
						errCh <- err
					}
				} else {
					b.logger.Warn("Unknown callback data: %s", callbackData)
				}

				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
				if _, err := b.API.Request(callback); err != nil {
					b.logger.Error("Ошибка при ответе на CallbackQuery: %v", err)
				}
			}
		}
	}()

	for {
		select {
		case err := <-errCh:
			b.logger.Error("Error from handler: %v", err)
		}
	}
}
