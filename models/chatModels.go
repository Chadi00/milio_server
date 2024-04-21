package models

import "time"

type ClaudeAPI struct {
	ID           string        `json:"id"`
	Type         string        `json:"type"`
	Role         string        `json:"role"`
	Model        string        `json:"model"`
	StopSequence *string       `json:"stop_sequence"`
	Usage        UsageClaude   `json:"usage"`
	Content      []ContentItem `json:"content"`
	StopReason   string        `json:"stop_reason"`
}

type UsageClaude struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type ContentItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type UserChat struct {
	ID      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type SystemChat struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
