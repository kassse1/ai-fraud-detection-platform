package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ScamMLRequest struct {
	Text string `json:"text"`
}

type ScamMLResponse struct {
	Score float64 `json:"score"`
}

type ScamMLClient struct {
	baseURL string
	client  *http.Client
}

func NewScamMLClient(baseURL string) *ScamMLClient {
	return &ScamMLClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *ScamMLClient) Analyze(ctx context.Context, text string) (*ScamMLResponse, error) {
	payload, err := json.Marshal(ScamMLRequest{Text: text})
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

	var result ScamMLResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
