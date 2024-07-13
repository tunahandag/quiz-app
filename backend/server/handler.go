package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetQuestions returns the list of questions as JSON
func GetQuestions(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Questions are requested!")
    json.NewEncoder(w).Encode(questions)
}

// Answer represents a single answer to a question
type Answer struct {
    QuestionID int `json:"question_id"`
    Answer     int `json:"answer"`
}

// Result represents the result of a quiz submission
type Result struct {
    CorrectAnswers int    `json:"correct_answers"`
    Comparison     string `json:"comparison"`
}

// submissions stores all quiz submissions
var submissions = [][]Answer{}

// SubmitAnswers receives quiz answers and calculates the result
func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
    var answers []Answer
    json.NewDecoder(r.Body).Decode(&answers)

    correct := 0
    for _, answer := range answers {
        for _, question := range questions {
            if question.ID == answer.QuestionID && question.Correct == answer.Answer {
                correct++
                break
            }
        }
    }

    submissions = append(submissions, answers)

    // Calculate comparison
    betterThan := 0
    // compare the current submission with all previous submissions
    for _, submission := range submissions {
        score := 0
        // calculate the score of the current submission
        for _, answer := range submission {
            for _, question := range questions {
                if question.ID == answer.QuestionID && question.Correct == answer.Answer {
                    score++
                    break
                }
            }
        }
        if score < correct {
            betterThan++
        }
    }
    // the first cannot be better than anyone
    comparison := "You were better than 0% of all quizzers"
    // if there are more than one submission, calculate the comparison
    if len(submissions) > 1 {
        comparisonPercentage := (betterThan * 100) / (len(submissions) - 1)
        comparison = "You were better than " + strconv.Itoa(comparisonPercentage) + "% of all quizzers"
    }

    json.NewEncoder(w).Encode(Result{CorrectAnswers: correct, Comparison: comparison})
}