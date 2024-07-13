package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
// It sends a GET request to /questions to get the questions
// It prints the questions and answers to the console
// It sends a POST request to /submit with the answers
// It prints the result of the quiz
var quizCmd = &cobra.Command{
    Use:   "take",
    Short: "Take the quiz",
    Run: func(cmd *cobra.Command, args []string) {
        // Get questions
        resp, err := http.Get("http://localhost:8080/questions")
        if err != nil {
            fmt.Println("Error getting questions:", err)
            return
        }
        // Close the response body when the function returns
        defer resp.Body.Close()

        fmt.Println("Welcome to the quiz!")
        fmt.Println("Please answer the following questions by picking the correct number for the answer.")
        
        // Read the response body
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response:", err)
            return
        }

        // Unmarshal the JSON response into a slice of questions    
        var questions []struct {
            ID       int      `json:"id"`
            Question string   `json:"question"`
            Answers  []string `json:"answers"`
        }
        json.Unmarshal(body, &questions)

        // Ask the questions and get the answers
        var answers []struct {
            QuestionID int `json:"question_id"`
            Answer     int `json:"answer"`
        }

        // Loop through the questions and answers
        for _, question := range questions {
            fmt.Println(question.Question)
            for i, answer := range question.Answers {
                fmt.Printf("%d. %s\n", i+1, answer)
            }
            var answer int
            fmt.Scan(&answer)
            answers = append(answers, struct {
                QuestionID int `json:"question_id"`
                Answer     int `json:"answer"`
            }{QuestionID: question.ID, Answer: answer - 1})
        }

        // Submit answers
        answerBytes, err := json.Marshal(answers)
        if err != nil {
            fmt.Println("Error marshalling answers:", err)
            return
        }

        // Send a POST request to /submit with the answers
        resp, err = http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(answerBytes))
        if err != nil {
            fmt.Println("Error submitting answers:", err)
            return
        }
        defer resp.Body.Close()

        // Read the response body
        body, err = io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response:", err)
            return
        }

        // Unmarshal the JSON response into a result struct
        var result struct {
            CorrectAnswers int    `json:"correct_answers"`
            Comparison     string `json:"comparison"`
        }
        json.Unmarshal(body, &result)
        
        fmt.Printf("You got %d correct answers.\n", result.CorrectAnswers)
        fmt.Println(result.Comparison)
    },
}

func init() {
    rootCmd.AddCommand(quizCmd)
}
