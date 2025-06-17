package quiz_test

import (
	"bytes"
	"io"
	"slices"
	"testing"
	"testing/fstest"
	"time"

	"github.com/bgics/gophercises/quiz/quiz"
)

func TestQuestionsFromCSV(t *testing.T) {
	const (
		csvFileName = "qna.csv"
		csvFileData = `5+5,  10
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

	output := bytes.Buffer{}

	t.Run("give answers before timer", func(t *testing.T) {
		input := DelayedReader{
			lines: []string{" 10", "10"},
			delay: 1 * time.Millisecond,
		}

		got := quiz.Quiz(questions, &input, &output, 5*time.Millisecond)

		want := 2

		if got != want {
			t.Errorf("got %d correct answers, want %d correct answers", got, want)
		}
	})
	t.Run("timer should run out", func(t *testing.T) {
		input := DelayedReader{
			lines: []string{"10", "10"},
			delay: 3 * time.Millisecond,
		}

		got := quiz.Quiz(questions, &input, &output, 5*time.Millisecond)

		want := 1

		if got != want {
			t.Errorf("got %d correct answers, want %d correct answers", got, want)
		}
	})
}

type DelayedReader struct {
	lines []string
	index int
	delay time.Duration
}

func (r *DelayedReader) Read(p []byte) (int, error) {
	if r.index >= len(r.lines) {
		return 0, io.EOF
	}

	time.Sleep(r.delay)

	line := r.lines[r.index] + "\n"
	n := copy(p, []byte(line))
	r.index += 1

	return n, nil
}
