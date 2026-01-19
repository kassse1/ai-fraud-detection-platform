package domain

import "time"

type AnalysisResultWithText struct {
	Text          string    `json:"text"`
	RiskScore     float64   `json:"risk_score"`
	IsScam        bool      `json:"is_scam"`
	IsAIGenerated bool      `json:"is_ai_generated"`
	Explanation   string    `json:"explanation"`
	CreatedAt     time.Time `json:"created_at"`
}
