package main

import (
	"log"
	"net/http"

	apihttp "github.com/kassse1/ai-fraud-backend/internal/http"
	"github.com/kassse1/ai-fraud-backend/internal/http/handlers"

	"github.com/kassse1/ai-fraud-backend/internal/db"
	"github.com/kassse1/ai-fraud-backend/internal/orchestrator"
	"github.com/kassse1/ai-fraud-backend/internal/repository"
)

func main() {
	// 🔹 PostgreSQL DSN (замени пароль)
	dsn := "postgres://postgres:040806@localhost:5432/ai_fraud?sslmode=disable"

	// 🔹 DB
	dbConn := db.NewPostgres(dsn)
	repo := repository.NewAnalysisRepository(dbConn)
	
	// 🔹 Orchestrator
	orch := orchestrator.NewOrchestrator(repo)

	// 🔹 HTTP handler
	analyzeHandler := handlers.NewAnalyzeHandler(orch)

	// 🔹 Router
	router := apihttp.NewRouter(analyzeHandler)

	log.Println("🚀 AI Fraud Detection API started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
