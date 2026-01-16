package http

import (
	"net/http"

	"github.com/kassse1/ai-fraud-backend/internal/http/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	analyzeHandler := handlers.NewAnalyzeHandler()

	mux.HandleFunc("/analyze", analyzeHandler.Handle)

	return mux
}
