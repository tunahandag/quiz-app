# Quiz Application

---

This project is a simple quiz application built with Golang. It includes a backend server that serves quiz questions and handles submissions, and a CLI client that allows users to take the quiz.

## User stories/Use cases:

- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"

## Features

- **Play Quiz:** Users can take a quiz, answer questions, and see their scores.
- **CLI Interface:** A command-line interface allows users to interact with the application.

## Technologies Used

- **Golang:** Backend and CLI implementation.
- **Cobra:** CLI framework for handling user inputs.

## Project Structure

```
quiz-app/
├── backend/
│   ├── main.go
│   ├── server/
│   │   ├── handler.go
│   │   └── questions.go
├── cli/
│   ├── main.go
│   └── cmd/
│       ├── root.go
│       └── quiz.go
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── questions.json
└── README.md
```

## Getting Started

### Prerequisites

- **Go (Golang):** Ensure you have Go installed. You can download it from [golang.org](https://golang.org/).

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/tunahandag/quiz-app.git
   cd quiz-app
   ```

2. Initialize Go modules:

   ```sh
   go mod init quiz-app
   ```

3. Install dependencies:

   ```sh
   go get github.com/spf13/cobra@latest
   ```

### Running the Application

#### 1. Backend Server

Start the backend server to serve quiz questions and handle submissions.

```sh
go run backend/main.go
```

#### 2. CLI Application

Start the CLI application to interact with the quiz.

```sh
go run cli/main.go take
```

### Usage

When you run the CLI application, you will start the quiz with the first question.

#### Playing the Quiz

- Answer the questions displayed.
- View your score and how you compared to other users.

## API Endpoints

The backend server exposes the following API endpoints:

- **GET /questions:** Retrieve all quiz questions.
- **POST /submit:** Submit quiz answers and receive the score.

## Example Questions (questions.json)

```json
[
  {
    "id": 1,
    "question": "What is the capital of France?",
    "answers": ["Berlin", "Madrid", "Paris", "Lisbon"],
    "correct": 2
  },
  {
    "id": 2,
    "question": "Which planet is known as the Red Planet?",
    "answers": ["Earth", "Mars", "Jupiter", "Venus"],
    "correct": 1
  },
  {
    "id": 3,
    "question": "What is the largest ocean on Earth?",
    "answers": ["Atlantic Ocean", "Indian Ocean", "Arctic Ocean", "Pacific Ocean"],
    "correct": 3
  },
  {
    "id": 4,
    "question": "Who wrote 'To Kill a Mockingbird'?",
    "answers": ["Harper Lee", "Mark Twain", "Ernest Hemingway", "F. Scott Fitzgerald"],
    "correct": 0
  },
  {
    "id": 5,
    "question": "What is the chemical symbol for gold?",
    "answers": ["Au", "Ag", "Pb", "Fe"],
    "correct": 0
  }
]
```

## Contributing

If you wish to contribute to this project, please fork the repository and submit a pull request. We welcome all contributions, including bug fixes, new features, and improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Golang](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

## Contact

For any questions or feedback, please contact:

- Tunahan Dağ: [tunahandag@hotmail.com](mailto:tunahandag@hotmail.com)
- GitHub: [tunahandag](https://github.com/tunahandag)
