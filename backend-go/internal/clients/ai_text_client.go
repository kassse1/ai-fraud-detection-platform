package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// ===== Request / Response =====

type AITextRequest struct {
	Text string `json:"text"`
}

type AITextResponse struct {
	IsAIGenerated bool    `json:"is_ai_generated"`
	Score         float64 `json:"score"`
	Explanation   string  `json:"explanation"`
}

// ===== Client =====

type AITextClient struct {
	baseURL string
	client  *http.Client
}

func NewAITextClient(baseURL string) *AITextClient {
	return &AITextClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// ===== Analyze =====

func (c *AITextClient) Analyze(ctx context.Context, text string) (*AITextResponse, error) {
	payload, err := json.Marshal(AITextRequest{Text: text})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+"/analyze",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AITextResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
