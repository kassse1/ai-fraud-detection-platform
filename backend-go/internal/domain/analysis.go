package domain

type AnalysisRequest struct {
	Text string `json:"text"`
}

type AnalysisResult struct {
	RiskScore     float64 `json:"risk_score"`
	IsScam        bool    `json:"is_scam"`
	IsAIGenerated bool    `json:"is_ai_generated"`
	Explanation   string  `json:"explanation"`
}
