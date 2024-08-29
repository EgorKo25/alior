package bot

func (b *Bot) Run() error {
	errCh := make(chan error, 1)

	go func() {
		updates := b.API.GetUpdatesChan(b.UpdateConfig)

		for update := range updates {
			if update.Message == nil || !update.Message.IsCommand() {
				continue
			}

			b.logger.Info("Got command: %s", update.Message.Text)
			command := update.Message.Command()

			if handler, exists := b.handlers[command]; exists {
				handler(&update)
			} else {
				b.unknownCommandHandler(&update)
			}
		}
	}()

	return <-errCh
}
