package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type CompletionResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
	Index   int     `json:"index"`
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: time.Second * 120},
		BaseURL:    baseURL,
		APIKey:     apiKey,
	}
}

func (c *Client) CreateCompletion(msg string, model string) (*CompletionResponse, error) {
	request := CompletionRequest{
		Model:    model,
		Messages: []Message{{Role: "user", Content: msg}},
	}

	resp, err := c.makeRequest(request)
	if err != nil {
		return nil, err
	}

	var result CompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &result, nil
}

func (c *Client) makeRequest(request CompletionRequest) (*http.Response, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/chat/completions", c.BaseURL), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Set("Content-Type", "application/json")

	return c.HTTPClient.Do(req)
}
