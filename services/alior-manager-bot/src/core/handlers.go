package bot

import (
	"context"
	"fmt"
	"github.com/EgorKo25/common/broker"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rabbitmq/amqp091-go"
)

func (b *Bot) initHandlers() {
	b.handlers["get_initial_callback"] = b.getInitialCallbackHandler
	//b.handlers["previous_callback"] = b.getInitialCallbackHandler
	b.handlers["next_callback"] = b.getNextCallbackHandler
	//b.handlers["delete_callback"] = b.getInitialCallbackHandler
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

	msg, ok := <-response
	if !ok {
		b.logger.Error("error getting initial callback response")
		return fmt.Errorf("error getting initial callback response")
	}
	b.logger.Info("Получено сообщение: %s", msg.Body)

	err = msg.Ack(false)
	if err != nil {
		b.logger.Error("Ошибка при подтверждении сообщения: %v", err)
		return err
	}

	message := tgbotapi.NewMessage(update.Message.Chat.ID, string(msg.Body))

	buttonPrev := tgbotapi.NewInlineKeyboardButtonData("Предыдущее", "previous_callback")
	buttonNext := tgbotapi.NewInlineKeyboardButtonData("Следующее", "next_callback")
	buttonDelete := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_callback")

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttonPrev, buttonNext),
		tgbotapi.NewInlineKeyboardRow(buttonDelete),
	)

	message.ReplyMarkup = keyboard

	_, err = b.API.Send(message)
	if err != nil {
		b.logger.Error("Ошибка при отправке сообщения: %v", err)
		return err
	}

	return nil
}

func (b *Bot) getNextCallbackHandler(ctx context.Context, update *tgbotapi.Update) error {
	b.logger.Info("in next msg handler")
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Следующее сообщение")
	_, err := b.API.Send(msg)
	if err != nil {
		b.logger.Error("Ошибка при отправке сообщения: %v", err)
		return err
	}

	return err
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
