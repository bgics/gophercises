package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bgics/gophercises/quiz/quiz"
)

var (
	csvFlag   = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitFlag = flag.Int("limit", 30, "the time limit for the quiz in seconds")
)

func main() {
	flag.Parse()

	questions, err := quiz.QuestionsFromCSV(os.DirFS("."), *csvFlag)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Press Enter to start...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	timeLimit := time.Duration(*limitFlag) * time.Second
	correctAnswers := quiz.Quiz(questions, os.Stdin, os.Stdout, timeLimit)

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(questions))
}
