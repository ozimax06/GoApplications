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

	//time.AfterFunc(4*time.Second, func() { fmt.Println("finished!") })
	quizProblems := loadProblemsFromCvs("problems.csv")
	correctCount, wrongCount := makeQuiz(quizProblems, 2)
	showResults(quizProblems, correctCount, wrongCount)

}

func makeQuiz(problems []Problem, seconds time.Duration) (int, int) {
	
	correctAnswers := 0
	wrongAnswers := 0

	time.AfterFunc(seconds*time.Second, func() {
		
		//return correctAnswers, wrongAnswers
	})

	for i := 0; i < len(problems); i++ {

		answer := AskQuestionAndGetAnswer(problems[i].Question)
		if answer == problems[i].Answer {
			correctAnswers++
		} else {
			wrongAnswers++
		}
	}

	return correctAnswers, wrongAnswers
}

func showResults(problems []Problem, correctAnswers int, wrongAnswers int) {
	fmt.Println("Result")
	fmt.Println("=============")

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
