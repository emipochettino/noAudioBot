package application

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"noAudioBot/internal/domain"
)

type MessageHandler struct {
	messageServices []domain.MessageService
}

func NewMessageHandler(commands ...domain.MessageService) *MessageHandler {
	return &MessageHandler{commands}
}

func (m MessageHandler) Handle(update *tgbotapi.Update) {
	for i := range m.messageServices {
		m.messageServices[i].Process(update)
	}
}
