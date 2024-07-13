package server

import (
	"encoding/json"
	"log"
	"os"
)

// Question represents a single question in the quiz
// It has an ID, a question, a list of answers and the index of the correct answer
type Question struct {
    ID       int      `json:"id"`
    Question string   `json:"question"`
    Answers  []string `json:"answers"`
    Correct  int      `json:"correct"`
}

// questions stores all questions
var questions []Question

// LoadQuestions reads questions from a file and stores them in the questions variable
// It logs a fatal error if the file cannot be read or the questions cannot be unmarshaled
func LoadQuestions(filename string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatalf("Failed to read questions file: %v", err)
    }
    err = json.Unmarshal(data, &questions)
    if err != nil {
        log.Fatalf("Failed to unmarshal questions: %v", err)
    }
}
