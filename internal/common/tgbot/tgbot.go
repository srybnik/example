package tgbot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBot struct {
	bot    *tgbotapi.BotAPI
	token  string
	chatID int64
}

func New(token string, chatID int64) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	t := TgBot{
		bot:    bot,
		token:  token,
		chatID: chatID,
	}
	return &t, nil
}

func (t *TgBot) Send(format string, a ...interface{}) error {
	text := fmt.Sprintf(format, a...)
	msg := tgbotapi.NewMessage(t.chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(msg)
	return err
}
