package main

import (
	"log"
	"net/http"

	internalhttp "github.com/kassse1/ai-fraud-backend/internal/http"
)

func main() {
	router := internalhttp.NewRouter()

	log.Println("🚀 AI Fraud Detection API started on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
