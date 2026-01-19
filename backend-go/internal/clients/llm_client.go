package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// ===== Request / Response =====

type LLMRequest struct {
	Text string `json:"text"`
}

type LLMResponse struct {
	Score       float64 `json:"score"`
	Explanation string  `json:"explanation"`
}

// ===== Client =====

type LLMClient struct {
	baseURL string
	client  *http.Client
}

func NewLLMClient(baseURL string) *LLMClient {
	return &LLMClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// ===== Analyze =====

func (c *LLMClient) Analyze(ctx context.Context, text string) (*LLMResponse, error) {
	payload, err := json.Marshal(LLMRequest{Text: text})
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

	var result LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
