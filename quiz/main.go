package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bgics/gophercises/quiz/quiz"
)

func main() {
	questions, err := quiz.QuestionsFromCSV(os.DirFS("."), "problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	correctAnswers := quiz.Quiz(questions, os.Stdin, os.Stdout)
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(questions))
}
