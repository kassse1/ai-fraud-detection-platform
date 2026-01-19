package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kassse1/ai-fraud-backend/internal/orchestrator"
)

type AnalyzeHandler struct {
	orch *orchestrator.Orchestrator
}

func NewAnalyzeHandler(orch *orchestrator.Orchestrator) *AnalyzeHandler {
	return &AnalyzeHandler{
		orch: orch,
	}
}

func (h *AnalyzeHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result := h.orch.Analyze(req.Text)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
