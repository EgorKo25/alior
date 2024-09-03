package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
}

func (b *Bot) getInitialCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	//TODO: все это вынести в отдельную функцию

	//TODO: broker.publish() -> content_type=callback, type=create
	//TODO: определить стурктуру сообщения rabbitmq(взять из cms)
	//TODO: msg := broker.consume()
	//TODO: return msg

	//TODO: конец функции

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "я пока в разработке") //TODO: в text отправить (returned msg)
	if _, err := b.API.Send(msg); err != nil {
		b.logger.Error("failed to send message: %s", err)
		return err
	}
	return nil
}

func (b *Bot) unknownCommandHandler(ctx context.Context, update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"я хз че ты вкинул мб самое время вытащить насвай из под губы и глянуть список команд")
	if _, err := b.API.Send(msg); err != nil {
		b.logger.Error("failed to send message: %s", err)
		return err
	}
	return nil
}
