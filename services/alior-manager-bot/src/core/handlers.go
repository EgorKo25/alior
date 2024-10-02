package bot

import (
	"alior-manager-bot/src/transport/broker"
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
}

func (b *Bot) getInitialCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	response, err := b.broker.CMSMessageExchange(ctx, "привет галава", "initial")
	if err != nil {
		b.logger.Error("error getting initial callback response", err)
		return err
	}

	err = json.Unmarshal(response.Body, &broker.Callback{})
	if err != nil {
		b.logger.Error("error unmarshalling initial callback response", err)
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(response.Body))
	if _, err := b.API.Send(msg); err != nil {
		b.logger.Error("failed to send message: %s", err)
		return err
	}
	return nil
}

func (b *Bot) unknownCommandHandler(update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"я хз че ты вкинул мб самое время вытащить насвай из под губы и глянуть список команд")
	if _, err := b.API.Send(msg); err != nil {
		b.logger.Error("failed to send message: %s", err)
		return err
	}
	return nil
}
