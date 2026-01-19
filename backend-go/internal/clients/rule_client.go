package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type RuleRequest struct {
	Text string `json:"text"`
}

type RuleResponse struct {
	Score       float64 `json:"score"`
	Explanation string  `json:"explanation"`
}

type RuleClient struct {
	baseURL string
	client  *http.Client
}

func NewRuleClient(baseURL string) *RuleClient {
	return &RuleClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *RuleClient) Analyze(ctx context.Context, text string) (*RuleResponse, error) {
	payload, err := json.Marshal(RuleRequest{Text: text})
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

	var result RuleResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
