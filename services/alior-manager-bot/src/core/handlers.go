package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
}

func (b *Bot) getInitialCallbackHandler(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "я пока в разработке")
	if _, err := b.API.Send(msg); err != nil {
		//b.logger.Fatal("Failed to send message: %s", err)
	}
}

func (b *Bot) unknownCommandHandler(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"я хз че ты вкинул мб самое время вытащить насвай из под губы и глянуть список команд")
	if _, err := b.API.Send(msg); err != nil {
		//b.logger.Fatal("Failed to send message: %s", err)
	}
}
