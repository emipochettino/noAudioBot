package main

import (
	"noAudioBot/internal/infrastructure"
)

func main() {
	app, err := infrastructure.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}
