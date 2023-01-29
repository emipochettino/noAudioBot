package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"noAudioBot/internal/domain"
	"time"
)

type AssemblyAIClient struct {
	apiKey string
	client *http.Client
}

func (a *AssemblyAIClient) Transcript(audioUrl string) string {
	values := map[string]string{"audio_url": audioUrl, "language_code": "es"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatalln(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.assemblyai.com/v2/transcript", bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", a.apiKey)
	res, err := a.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	// TODO check if it can return an error

	return result["id"].(string)
}

func (a *AssemblyAIClient) FindTranscription(id string) string {
	var transcription string
	for {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.assemblyai.com/v2/transcript/%s", id), nil)
		req.Header.Set("content-type", "application/json")
		req.Header.Set("authorization", a.apiKey)
		res, err := a.client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)

		if result["status"].(string) == "processing" {
			time.Sleep(1 * time.Second)
			continue
		} else {
			transcription = fmt.Sprintf("%s\n", result["text"])
			break
		}
	}

	return transcription
}

func NewAssemblyAIClient(apiKey string, client *http.Client) domain.TranscriptionRepository {
	return &AssemblyAIClient{apiKey: apiKey, client: client}
}
