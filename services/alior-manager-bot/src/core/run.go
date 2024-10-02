package bot

import (
	"context"
)

func (b *Bot) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	defer close(errCh)

	go func() {
		updates := b.API.GetUpdatesChan(b.UpdateConfig)

		for update := range updates {
			if update.Message == nil || !update.Message.IsCommand() {
				continue
			}

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
	}()

	for {
		select {
		case err := <-errCh:
			b.logger.Error("Error from handler: %v", err)
		}
	}
}
