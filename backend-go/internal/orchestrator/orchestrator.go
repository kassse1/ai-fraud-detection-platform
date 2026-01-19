package orchestrator

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/kassse1/ai-fraud-backend/internal/clients"
	"github.com/kassse1/ai-fraud-backend/internal/domain"
	"github.com/kassse1/ai-fraud-backend/internal/repository"
)

type Orchestrator struct {
	repo         *repository.AnalysisRepository
	ruleClient   *clients.RuleClient
	scamMLClient *clients.ScamMLClient
	llmClient    *clients.LLMClient
	aiTextClient *clients.AITextClient
}

func NewOrchestrator(repo *repository.AnalysisRepository) *Orchestrator {
	return &Orchestrator{
		repo:         repo,
		ruleClient:   clients.NewRuleClient("http://localhost:8001"),
		scamMLClient: clients.NewScamMLClient("http://localhost:8002"),
		llmClient:    clients.NewLLMClient("http://localhost:8003"),
		aiTextClient: clients.NewAITextClient("http://localhost:8004"),
	}
}

func (o *Orchestrator) Analyze(text string) domain.AnalysisResult {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	var (
		ruleScore   float64
		mlScore     float64
		llmScore    float64
		isAIGen     bool
		explanation string
	)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		res, err := o.ruleClient.Analyze(ctx, text)
		if err != nil {
			log.Println("rule engine failed:", err)
			return
		}
		ruleScore = res.Score
		explanation = res.Explanation
	}()
	go func() {
		defer wg.Done()
		res, err := o.llmClient.Analyze(ctx, text)
		if err != nil {
			log.Println("llm analyzer failed:", err)
			return
		}
		llmScore = res.Score
		if explanation == "" {
			explanation = res.Explanation
		}
	}()

	go func() {
		defer wg.Done()
		res, err := o.scamMLClient.Analyze(ctx, text)
		if err != nil {
			log.Println("scam-ml failed:", err)
			return
		}
		mlScore = res.Score
	}()
	go func() {
		defer wg.Done()
		res, err := o.aiTextClient.Analyze(ctx, text)
		if err != nil {
			log.Println("ai-text-detector failed:", err)
			return
		}
		isAIGen = res.IsAIGenerated
		if explanation == "" {
			explanation = res.Explanation
		}
	}()

	wg.Wait()

	finalRisk := aggregateRisk(ruleScore, mlScore, llmScore)

	result := domain.AnalysisResult{
		RiskScore:     finalRisk,
		IsScam:        finalRisk > 0.6,
		IsAIGenerated: isAIGen,
		Explanation:   explanation,
	}

	// 🔥 ВОТ ЭТОГО СЕЙЧАС НЕ ХВАТАЕТ
	if err := o.repo.Save(ctx, result, text); err != nil {
		log.Println("❌ failed to save analysis:", err)
	} else {
		log.Println("✅ analysis saved to DB")
	}

	return result
}
func (o *Orchestrator) GetHistory(
	ctx context.Context,
	limit int,
	offset int,
) ([]domain.AnalysisResultWithText, error) {
	return o.repo.List(ctx, limit, offset)
}

func aggregateRisk(rule, ml, llm float64) float64 {
	const (
		ruleWeight = 0.2
		mlWeight   = 0.2
		llmWeight  = 0.6
	)
	return rule*ruleWeight + ml*mlWeight + llm*llmWeight
}
