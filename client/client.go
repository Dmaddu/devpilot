package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GPTClient represents a client for communicating with GPT-4
type GPTClient struct {
	apiKey     string
	apiURL     string
	httpClient *http.Client
}

// Message represents a role-based message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GPTRequest represents the payload sent to the GPT API
type GPTRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

// Choice represents one of the possible responses from GPT-4
type Choice struct {
	Message Message `json:"message"`
}

// GPTResponse represents the response from the GPT API
type GPTResponse struct {
	Choices []Choice `json:"choices"`
}

func NewAzureOpenAIClient() *GPTClient {
	azureEndpoint := "https://in-ep-one.openai.azure.com"
	apiKey := "FEHb9gIIByqzOKz0mczL0rU4heu5veqLu6OMSKuHHVRWMeFR80NlJQQJ99BCAC77bzfXJ3w3AAABACOGQdFM"
	deploymentName := "gpt-4o"
	apiVersion := "2024-05-01-preview"

	endpoint := strings.TrimRight(azureEndpoint, "/")
	apiURL := fmt.Sprintf("%s/openai/deployments/%s/chat/completions?api-version=%s", endpoint, deploymentName, apiVersion)
	return &GPTClient{
		apiKey:     apiKey,
		apiURL:     apiURL,
		httpClient: &http.Client{},
	}
}

// SendPrompt sends a prompt to GPT-4 and returns the response text.
func (c *GPTClient) SendPrompt(prompt string) (string, error) {
	if c.apiKey == "" {
		return "", errors.New("OPENAI_API_KEY is not set")
	}

	requestPayload := GPTRequest{
		Model:    "gpt-4",
		Messages: []Message{{Role: "user", Content: prompt}},
	}

	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse GPTResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		return "", err
	}

	if len(apiResponse.Choices) == 0 {
		return "", errors.New("no choices found in response: " + string(bodyBytes))
	}
	return apiResponse.Choices[0].Message.Content, nil
}
