package infrastructure

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Bot struct {
	BotApi       *telegram.BotAPI
	allowedUsers []int64
}

func NewBot(telegramToken, profile string, allowedUsers []int64) (*Bot, error) {
	botApi, err := telegram.NewBotAPI(telegramToken)
	if err != nil {
		return nil, err
	}
	botApi.Debug = strings.EqualFold("dev", profile)

	return &Bot{botApi, allowedUsers}, nil
}

func (bot *Bot) FindUrlById(id string) (string, error) {
	return bot.BotApi.GetFileDirectURL(id)
}

func (bot *Bot) Create(messageConfig telegram.MessageConfig) error {
	_, err := bot.BotApi.Send(messageConfig)
	return err
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
