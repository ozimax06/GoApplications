package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	//go run main.go --help
	//go run main.go -duration 20
	quizDuration := flag.Int("duration", 10, "Duration Seconds?")
	flag.Parse()
	startQuiz(*quizDuration)

}

func startQuiz(duration int) {
	QuizDurationSeconds := time.Duration(duration)
	quizChanel := make(chan bool, 1)
	correctAnswers := 0
	wrongAnswers := 0
	quizProblems := loadProblemsFromCvs("problems.csv")

	// Run Quiz in it's own goroutine and pass back it's response into our channel.
	go func() {
		result := startAskingQuizQuestions(quizProblems, &correctAnswers, &wrongAnswers)
		quizChanel <- result
	}()

	// Listen on our channel AND a timeout channel - which ever happens first.
	select {
	case <-quizChanel:
		displayQuizFinishedMessage()
	case <-time.After(QuizDurationSeconds * time.Second):
		displayTimesUpMessage()
	}

	showResults(quizProblems, correctAnswers, wrongAnswers)
}

func startAskingQuizQuestions(problems []Problem, correctAnswers *int, wrongAnswers *int) bool {

	for i := 0; i < len(problems); i++ {
		answer := askQuestionAndGetAnswer(problems[i].Question)
		if answer == problems[i].Answer {
			*correctAnswers++
		} else {
			*wrongAnswers++
		}
	}
	return true
}

func showResults(problems []Problem, correctAnswers int, wrongAnswers int) {
	fmt.Println("RESULT:")
	fmt.Println("Correct Answers: ", correctAnswers)
	fmt.Println("Wrong Answers: ", wrongAnswers)
	fmt.Println("Total: ", len(problems))
}

func askQuestionAndGetAnswer(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question, " ?")
	fmt.Println("-> ")
	answer, _ := reader.ReadString('\n')
	answer = strings.Replace(answer, "\n", "", -1)

	return answer
}

func loadProblemsFromCvs(filename string) []Problem {
	f, _ := os.Open(filename)

	r := csv.NewReader(bufio.NewReader(f))
	problems := []Problem{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		problem := Problem{Question: record[0], Answer: record[1]}
		problems = append(problems, problem)
	}
	return problems
}

func displayQuizFinishedMessage() {
	fmt.Println("=========")
	fmt.Println("YOU FINISHED THE QUIZ!")
	fmt.Println("=========")
}

func displayTimesUpMessage() {
	fmt.Println("=========")
	fmt.Println("TIME'S UP!")
	fmt.Println("=========")
}

type Problem struct {
	Question string
	Answer   string
}
