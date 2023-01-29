package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TranscriptionRepository interface {
	Transcript(url string) string
	FindTranscription(id string) string
}

type FileRepository interface {
	FindUrlById(id string) (string, error)
}

type MessageRepository interface {
	Create(messageConfig tgbotapi.MessageConfig) error
}
