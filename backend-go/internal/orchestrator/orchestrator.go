package orchestrator

import "github.com/kassse1/ai-fraud-backend/internal/domain"

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) Analyze(text string) domain.AnalysisResult {
	// 🔴 ВАЖНО:
	// Сейчас mock.
	// Потом здесь будут параллельные вызовы Python AI сервисов

	return domain.AnalysisResult{
		RiskScore:     0.85,
		IsScam:        true,
		IsAIGenerated: false,
		Explanation:   "Mock response: AI services not connected yet",
	}
}
