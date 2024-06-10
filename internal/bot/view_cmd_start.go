package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ViewCmdStart(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	if _, err := bot.Send(tgbotapi.NewMessage(update.FromChat().ID, "Hello, I'm anekdot bot. I can send you anekdots.")); err != nil {
		return err
	}

	return nil
}
