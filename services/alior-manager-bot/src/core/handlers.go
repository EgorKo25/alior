package bot

import (
	"context"
	"fmt"
	"github.com/EgorKo25/common/broker"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
	b.handlers["previous_callback"] = b.getPreviousCallbackHandler
	b.handlers["next_callback"] = b.getNextCallbackHandler
	b.handlers["delete_callback"] = b.deleteCallbackHandler
}

func (b *Bot) getInitialCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	err := broker.Publish("ask_publisher",
		amqp091.Publishing{
			ContentType: "text/plain",
			Type:        "callback",
			Headers:     amqp091.Table{"action_type": "initial"},
		},
	)
	if err != nil {
		b.logger.Error("error getting initial callback response", err)
		return err
	}

	response, err := broker.Consume("ans_consumer")
	if err != nil {
		b.logger.Error("error consuming initial callback response", err)
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	select {
	case msg, ok := <-response:
		if !ok {
			b.logger.Error("error getting initial callback response: channel closed")
			return fmt.Errorf("error getting initial callback response: channel closed")
		}
		b.logger.Info("Получено сообщение: %s", msg.Body)

		if err := msg.Ack(false); err != nil {
			b.logger.Error("Ошибка при подтверждении сообщения: %v", err)
			return err
		}

		message := tgbotapi.NewMessage(update.Message.Chat.ID, string(msg.Body))
		message.ReplyMarkup = createKeyboardMarkup()

		if _, err := b.API.Send(message); err != nil {
			b.logger.Error("Ошибка при отправке сообщения: %v", err)
			return err
		}

	case <-ctx.Done():
		b.logger.Error("Timeout waiting for initial callback response")
		return ctx.Err()
	}
	return nil
}

func (b *Bot) getNextCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	err := broker.Publish("ask_publisher",
		amqp091.Publishing{
			ContentType: "text/plain",
			Type:        "callback",
			Headers:     amqp091.Table{"action_type": "next"},
		},
	)
	if err != nil {
		b.logger.Error("error getting next callback response", err)
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	response, err := broker.Consume("ans_consumer")
	if err != nil {
		b.logger.Error("error consuming next callback response", err)
		return err
	}

	select {
	case msg, ok := <-response:
		if !ok {
			b.logger.Error("error getting next callback response: channel closed")
			return fmt.Errorf("error getting next callback response: channel closed")
		}
		b.logger.Info("Получено сообщение: %s", msg.Body)

		if err := msg.Ack(false); err != nil {
			b.logger.Error("Ошибка при подтверждении сообщения: %v", err)
			return err
		}

		edit := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, string(msg.Body))
		edit.ReplyMarkup = createKeyboardMarkup()

		if _, err := b.API.Send(edit); err != nil {
			b.logger.Error("Ошибка при редактировании сообщения: %v", err)
			return err
		}

	case <-ctx.Done():
		b.logger.Error("Timeout waiting for next callback response")
		return ctx.Err()
	}

	return nil
}

func (b *Bot) getPreviousCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	b.logger.Info("in previous msg handler")

	edit := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "Предыдущее сообщение")

	buttonPrev := tgbotapi.NewInlineKeyboardButtonData("Предыдущее", "previous_callback")
	buttonNext := tgbotapi.NewInlineKeyboardButtonData("Следующее", "next_callback")
	buttonDelete := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_callback")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttonPrev, buttonNext),
		tgbotapi.NewInlineKeyboardRow(buttonDelete),
	)

	edit.ReplyMarkup = &keyboard

	_, err := b.API.Send(edit)
	if err != nil {
		b.logger.Error("Ошибка при редактировании сообщения: %v", err)
		return err
	}

	return nil
}

func (b *Bot) deleteCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	b.logger.Info("in delete msg handler")

	edit := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, "Сообщение удалено")

	buttonPrev := tgbotapi.NewInlineKeyboardButtonData("Предыдущее", "previous_callback")
	buttonNext := tgbotapi.NewInlineKeyboardButtonData("Следующее", "next_callback")
	buttonDelete := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_callback")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttonPrev, buttonNext),
		tgbotapi.NewInlineKeyboardRow(buttonDelete),
	)

	edit.ReplyMarkup = &keyboard

	_, err := b.API.Send(edit)
	if err != nil {
		b.logger.Error("Ошибка при редактировании сообщения: %v", err)
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

func createKeyboardMarkup() *tgbotapi.InlineKeyboardMarkup {
	buttonPrev := tgbotapi.NewInlineKeyboardButtonData("Предыдущее", "previous_callback")
	buttonNext := tgbotapi.NewInlineKeyboardButtonData("Следующее", "next_callback")
	buttonDelete := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_callback")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttonPrev, buttonNext),
		tgbotapi.NewInlineKeyboardRow(buttonDelete),
	)

	return &keyboard
}
