package apicalls

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"milio/models"
	"os"

	"net/http"

	"github.com/joho/godotenv"
)

func CallMistralAPI(message string) (*models.SystemChat, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("MIXTRAL_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not set in .env file")
	}

	url := "https://api.mistral.ai/v1/chat/completions" // Mistral API URL

	// Construct the request body
	requestBody := map[string]interface{}{
		"model": "open-mixtral-8x7b",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
		"temperature": 0.2,
		"top_p":       1,
		"max_tokens":  512,
		"stream":      false,
		"safe_prompt": false,
		"random_seed": 1337,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var chat models.SystemChat
	err = json.Unmarshal(body, &chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}
