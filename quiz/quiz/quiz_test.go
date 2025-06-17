package quiz_test

import (
	"bytes"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/bgics/gophercises/quiz/quiz"
)

func TestQuestionsFromCSV(t *testing.T) {
	const (
		csvFileName = "qna.csv"
		csvFileData = `5+5,10
7+3,10
`
	)

	fs := fstest.MapFS{
		csvFileName: {Data: []byte(csvFileData)},
	}

	want := []quiz.Question{
		{"5+5", "10"},
		{"7+3", "10"},
	}

	got, err := quiz.QuestionsFromCSV(fs, csvFileName)

	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestQuiz(t *testing.T) {
	questions := []quiz.Question{
		{"5+5", "10"},
		{"7+3", "10"},
	}

	input := `10
10
`
	buf := bytes.Buffer{}
	r := strings.NewReader(input)

	got := quiz.Quiz(questions, r, &buf)
	want := 2

	if got != want {
		t.Errorf("got %d correct answers, want %d correct answers", got, want)
	}
}

func TestTimedQuiz(t *testing.T) {

}
