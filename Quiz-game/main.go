package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	QuizDurationSeconds := time.Duration(5)
	quizChanel := make(chan bool, 1)
	correctAnswers := 0
	wrongAnswers := 0
	quizProblems := loadProblemsFromCvs("problems.csv")

	// Run Quiz in it's own goroutine and pass back it's response into our channel.
	go func() {
		result := makeQuiz(quizProblems, &correctAnswers, &wrongAnswers)
		quizChanel <- result
	}()

	// Listen on our channel AND a timeout channel - which ever happens first.
	select {
	case <-quizChanel:
		fmt.Println("=========")
		fmt.Println("YOU FINISHED THE QUIZ!")
		fmt.Println("=========")
	case <-time.After(QuizDurationSeconds * time.Second):
		fmt.Println("=========")
		fmt.Println("TIME'S UP!")
		fmt.Println("=========")
	}

	showResults(quizProblems, correctAnswers, wrongAnswers)

}

func makeQuiz(problems []Problem, correctAnswers *int, wrongAnswers *int) bool {
	
	for i := 0; i < len(problems); i++ {

		answer: =AskQuestionAndGetAnswer(problems[i].Question)
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

func AskQuestionAndGetAnswer(question string) string {
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

type Problem struct {
	Question string
	Answer   string
}
