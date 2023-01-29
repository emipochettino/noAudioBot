package infrastructure

import (
	"errors"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"noAudioBot/internal/application"
	"noAudioBot/internal/domain"
	"os"
	"strconv"
	"strings"
)

type App struct {
	bot            *Bot
	messageHandler *application.MessageHandler
}

func NewApp() (*App, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if len(telegramToken) == 0 {
		return nil, errors.New("TELEGRAM_TOKEN is not set")
	}
	assemblyAiApiKey := os.Getenv("ASSEMBLY_AI_API_KEY")
	if len(assemblyAiApiKey) == 0 {
		return nil, errors.New("ASSEMBLY_AI_API_KEY is not set")
	}
	allowedTelegramUsers := os.Getenv("ALLOWED_TELEGRAM_USERS")
	if len(allowedTelegramUsers) == 0 {
		return nil, errors.New("ALLOWED_TELEGRAM_USERS is not set")
	}
	allowedUsersIds := mapStringArrayToInt64Array(strings.Split(allowedTelegramUsers, ","))
	profile := os.Getenv("PROFILE")

	bot, err := NewBot(telegramToken, profile, allowedUsersIds)
	if err != nil {
		panic(err)
	}

	transcriptionRepository := NewAssemblyAIClient(assemblyAiApiKey, &http.Client{})

	audioMessageService := domain.NewAudioMessageService(transcriptionRepository, bot, bot)
	messageHandler := application.NewMessageHandler(audioMessageService)

	return &App{
		bot:            bot,
		messageHandler: messageHandler,
	}, nil
}

func (app *App) Run() {
	updateConfig := telegram.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := app.bot.BotApi.GetUpdatesChan(updateConfig)
	for update := range updates {
		if len(app.bot.allowedUsers) > 0 && contains(app.bot.allowedUsers, update.SentFrom().ID) {
			app.messageHandler.Handle(&update)
		}
	}
}

func mapStringArrayToInt64Array(array []string) []int64 {
	int64Array := make([]int64, 0)
	for i := range array {
		int64Array = append(int64Array, mapStringToInt64(array[i]))
	}

	return int64Array
}

func mapStringToInt64(number string) int64 {
	i, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}
