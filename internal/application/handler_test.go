package application

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"noAudioBot/internal/domain"
	"testing"
)

func TestHandle(t *testing.T) {
	t.Run(`	Given a list of mock services
					When an update is handle
					Then all services are executed`,
		func(t *testing.T) {
			messageServiceOneMock := &domain.MessageServiceMock{
				ProcessFunc: func(message *tgbotapi.Update) {
					assert.NotNil(t, message)
				},
			}
			messageServiceTwoMock := &domain.MessageServiceMock{
				ProcessFunc: func(message *tgbotapi.Update) {
					assert.NotNil(t, message)
				},
			}

			handler := NewMessageHandler(messageServiceOneMock, messageServiceTwoMock)

			handler.handle(&tgbotapi.Update{})
		})
}
