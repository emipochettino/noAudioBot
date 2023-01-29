package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

//go:generate moq -out message_service_mock.go . MessageService
type MessageService interface {
	Process(update *tgbotapi.Update)
}

type AudioMessageService struct {
	transcriptionRepository TranscriptionRepository
	fileRepository          FileRepository
	messageRepository       MessageRepository
}

func (a AudioMessageService) Process(update *tgbotapi.Update) {
	if update.Message.Voice != nil {
		url, err := a.fileRepository.FindUrlById(update.Message.Voice.FileID)
		if err != nil {
			err = a.messageRepository.Create(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			if err != nil {
				log.Println(err)
			}
		}
		transcriptionId := a.transcriptionRepository.Transcript(url)
		transcription := a.transcriptionRepository.FindTranscription(transcriptionId)
		err = a.messageRepository.Create(tgbotapi.NewMessage(update.Message.Chat.ID, transcription)) // send message with transcription
		if err != nil {
			log.Println(err)
		}
	}
}

func NewAudioMessageService(
	transcriptionRepository TranscriptionRepository,
	fileRepository FileRepository,
	messageRepository MessageRepository,
) MessageService {
	return &AudioMessageService{
		transcriptionRepository: transcriptionRepository,
		fileRepository:          fileRepository,
		messageRepository:       messageRepository,
	}
}
