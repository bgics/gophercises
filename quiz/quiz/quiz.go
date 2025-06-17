package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"time"
)

type Question struct {
	Problem  string
	Solution string
}

func Quiz(questions []Question, in io.Reader, out io.Writer, timeLimit time.Duration) int {
	scanner := bufio.NewScanner(in)
	timer := time.NewTimer(timeLimit)

	correct := 0
	for i, q := range questions {
		fmt.Fprintf(out, "Problem #%d: %s = ", i+1, q.Problem)

		answerChan := make(chan string, 1)
		go func() {
			scanner.Scan()
			answerChan <- scanner.Text()
		}()

		select {
		case <-timer.C:
			return correct
		case answer := <-answerChan:
			if strings.TrimSpace(answer) == strings.TrimSpace(q.Solution) {
				correct += 1
			}
		}
	}

	timer.Stop()
	return correct
}

func QuestionsFromCSV(fileSystem fs.FS, fileName string) ([]Question, error) {
	records, err := readCSVFile(fileSystem, fileName)

	if err != nil {
		return nil, err
	}

	return parseQuestions(records)
}

func readCSVFile(fileSystem fs.FS, fileName string) ([][]string, error) {
	file, err := fileSystem.Open(fileName)

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	return r.ReadAll()
}

func parseQuestions(records [][]string) ([]Question, error) {
	var questions []Question

	for i, row := range records {
		if len(row) != 2 {
			return nil, fmt.Errorf("incorrect csv row, [row %d] expected 2 elements, got %d elements", i, len(row))
		}

		q := Question{
			Problem:  row[0],
			Solution: strings.TrimSpace(row[1]),
		}
		questions = append(questions, q)
	}

	return questions, nil
}
