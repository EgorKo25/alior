package bot

import (
	"context"
	"errors"
	"github.com/EgorKo25/common/broker"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rabbitmq/amqp091-go"
)

const (
	SUCCESS = "success"
	ERROR   = "error"
)

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
	b.handlers["previous_callback"] = b.getPreviousCallbackHandler
	b.handlers["next_callback"] = b.getNextCallbackHandler
	b.handlers["delete_callback"] = b.deleteCallbackHandler
}

func (b *Bot) getInitialCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	msg, err := b.processCallbackMessage("initial")
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, string(msg.Body))
	message.ReplyMarkup = b.addKeyboard()

	_, err = b.API.Send(message)
	if err != nil {
		b.logger.Error("Ошибка при отправке сообщения: %v", err)
		return err
	}

	return nil
}

func (b *Bot) getNextCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	msg, err := b.processCallbackMessage("next")
	if err != nil {
		return err
	}

	edit := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, string(msg.Body))
	edit.ReplyMarkup = b.addKeyboard()

	_, err = b.API.Send(edit)
	if err != nil {
		b.logger.Error("Ошибка при редактировании сообщения: %v", err)
		return err
	}

	return nil
}

func (b *Bot) getPreviousCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	msg, err := b.processCallbackMessage("previous")
	if err != nil {
		return err
	}

	edit := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, string(msg.Body))
	edit.ReplyMarkup = b.addKeyboard()

	_, err = b.API.Send(edit)
	if err != nil {
		b.logger.Error("Ошибка при редактировании сообщения: %v", err)
		return err
	}

	return nil
}

func (b *Bot) deleteCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	msg, err := b.processCallbackMessage("delete")
	if err != nil {
		return err
	}

	deleteMessage := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, string(msg.Body))
	_, err = b.API.Send(deleteMessage)
	if err != nil {
		b.logger.Error("Ошибка при удалении сообщения: %v", err)
		return err
	}

	b.logger.Info("Сообщение успешно удалено: %s", msg.Body)

	err = b.getNextCallbackHandler(ctx, update)
	if err != nil {
		b.logger.Error("Ошибка при вызове next callback handler: %v", err)
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

func (b *Bot) processCallbackMessage(actionType string) (*amqp091.Delivery, error) {
	err := broker.Publish(b.PublisherName, amqp091.Publishing{
		ContentType: "text/plain",
		Type:        "callback",
		Headers:     amqp091.Table{"action_type": actionType},
	})
	if err != nil {
		b.logger.Error("error publishing callback request", err)
		return nil, err
	}

	response, err := broker.Consume(b.ConsumerName)
	if err != nil {
		b.logger.Error("error consuming callback response", err)
		return nil, err
	}

	msg, ok := <-response
	if !ok {
		b.logger.Error("error getting callback response")
		return nil, errors.New("error getting callback response")
	}

	if _, ok := msg.Headers["msg_type"]; !ok {
		b.logger.Error("invalid message header")
		return nil, errors.New("invalid message header")
	}

	switch msg.Headers["msg_type"] {
	case SUCCESS:
		b.logger.Info("Получено сообщение: %s", msg.Body)
	case ERROR:
		b.logger.Error("ошибка при получении сообщения: %s", msg.Body)
		return nil, errors.New("error getting callback response")
	}

	err = msg.Ack(false)
	if err != nil {
		b.logger.Error("Ошибка при подтверждении сообщения: %v", err)
		return nil, err
	}

	return &msg, nil
}

func (b *Bot) addKeyboard() *tgbotapi.InlineKeyboardMarkup {
	buttonPrev := tgbotapi.NewInlineKeyboardButtonData("Предыдущее", "previous_callback")
	buttonNext := tgbotapi.NewInlineKeyboardButtonData("Следующее", "next_callback")
	buttonDelete := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_callback")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttonPrev, buttonNext),
		tgbotapi.NewInlineKeyboardRow(buttonDelete),
	)

	return &keyboard
}
