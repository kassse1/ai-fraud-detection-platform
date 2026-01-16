package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kassse1/ai-fraud-backend/internal/domain"
	"github.com/kassse1/ai-fraud-backend/internal/orchestrator"
)

type AnalyzeHandler struct {
	orchestrator *orchestrator.Orchestrator
}

func NewAnalyzeHandler() *AnalyzeHandler {
	return &AnalyzeHandler{
		orchestrator: orchestrator.NewOrchestrator(),
	}
}

func (h *AnalyzeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req domain.AnalysisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := h.orchestrator.Analyze(req.Text)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
