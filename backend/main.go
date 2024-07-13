package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tunahandag/quiz-app/backend/server"
)

// Load questions from questions.jso file
// Run HTTP server on port 8080
// Handle /questions and /submit routes
func main() {
    server.LoadQuestions("questions.json")
    http.HandleFunc("/questions", server.GetQuestions)
    http.HandleFunc("/submit", server.SubmitAnswers)
    fmt.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
