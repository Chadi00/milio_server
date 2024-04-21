package apicalls

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"milio/models"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LLM_API(message string, maxToken int, temperature float64) (*models.SystemChat, error) {
	res, err := CallGroqAPI(message, maxToken, temperature)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CallMistralAPI(message string, maxToken int, temperature float64) (*models.SystemChat, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	apiKey := os.Getenv("MIXTRAL_API_KEY")
	if apiKey == "" {
		log.Fatal("API key for Mistral not set as env variable")
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
		"temperature": temperature,
		"top_p":       1,
		"max_tokens":  maxToken,
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

func CallOpenAIAPI(message string, maxToken int, temperature float64) (*models.SystemChat, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API key for Open AI not set as env variable")
	}

	url := "https://api.openai.com/v1/chat/completions" // OPENAI API URL

	// Construct the request body
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo-0125",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
		"temperature": temperature,
		"max_tokens":  maxToken,
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

func CallAnthropicAPI(message string, maxTokens int, temperature float64) (*models.SystemChat, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("Anthropic API key not set as env variable")
	}

	url := "https://api.anthropic.com/v1/messages"

	// Construct the request body
	requestBody := map[string]interface{}{
		"model":       "claude-3-haiku-20240307",
		"max_tokens":  maxTokens,
		"temperature": temperature,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
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
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response body
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

func streamHandler(ctx *gin.Context, resp *http.Response) error {
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)

	// Process the stream
	ctx.Stream(func(w io.Writer) bool {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from stream (non-EOF): %v", err)
			} else {
				log.Println("EOF reached in stream")
			}
			return false
		}

		if bytes.HasPrefix(line, []byte("data: ")) {
			trimmedLine := bytes.TrimSpace(line[6:])

			var event map[string]interface{}
			if err := json.Unmarshal(trimmedLine, &event); err != nil {
				log.Printf("Error unmarshalling event: %v", err)
				return true
			}

			eventType, ok := event["type"].(string)
			if !ok {
				log.Println("No event type found in event data")
				return true
			}

			switch eventType {
			case "content_block_delta":
				if delta, ok := event["delta"].(map[string]interface{}); ok {
					if text, ok := delta["text"].(string); ok {
						log.Printf("Sending content block delta text: %s", text)
						ctx.SSEvent("message", text)
					} else {
						log.Println("No text found in content block delta")
					}
				} else {
					log.Println("Malformed delta in content block event")
				}
			case "message_stop":
				log.Println("Message stop event received, ending stream")
				ctx.SSEvent("end", "true")
				return false
			default:
				log.Printf("Unhandled event type '%s' received", eventType)
				ctx.SSEvent(eventType, event)
			}
		}
		return true
	})

	log.Println("Stream handler completed")
	return nil
}

func CallAnthropicAPIStreaming(ctx *gin.Context, message string, maxTokens int, temperature float64) error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("Anthropic API key not set as env variable")
	}

	url := "https://api.anthropic.com/v1/messages"

	// Construct the request body
	requestBody := map[string]interface{}{
		"model":       "claude-3-haiku-20240307",
		"max_tokens":  maxTokens,
		"temperature": temperature,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
		"stream": true,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("anthropic-beta", "messages-2023-12-15")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to API: %v", err)
		return err
	}

	// Handle the streaming response
	return streamHandler(ctx, resp)
}

func CallGroqAPI(message string, maxToken int, temperature float64) (*models.SystemChat, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		log.Fatal("API key for GROQ not set as env variable")
	}

	url := "https://api.groq.com/openai/v1/chat/completions"

	// construct the request body
	requestBody := map[string]interface{}{
		"model": "llama3-70b-8192",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": message,
			},
		},
		"temperature": temperature,
		"max_tokens":  maxToken,
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

	// read the response body
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
