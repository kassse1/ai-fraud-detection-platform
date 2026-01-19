package http

import (
	"github.com/gorilla/mux"
	"github.com/kassse1/ai-fraud-backend/internal/http/handlers"
)

func NewRouter(analyzeHandler *handlers.AnalyzeHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/analyze", analyzeHandler.Analyze).Methods("POST")
	r.HandleFunc("/history", analyzeHandler.History).Methods("GET")
	return r
}
