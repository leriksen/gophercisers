package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const OK = 0

type problem struct {
	question string
	answer   string
}

type results struct {
	asked int
	right int
}

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout

func main() {

	filename := flag.String("filename", "./problems.csv", "filename to use if passed, ./problems.csv")
	flag.Parse()

	csvfile, err := openFile(*filename)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := csvfile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	problems, err := readProblems(csvfile)

	if err != nil {
		log.Fatal(err)
	}

	result := askQuestions(problems, out, in)

	result.show(out)
}

func toProblem(row []string) problem {
	return problem{
		strings.TrimSpace(row[0]),
		strings.TrimSpace(row[1]),
	}
}

func openFile(filename string) (csvfile io.ReadCloser, err error) {
	csvfile, err = os.Open(filename)

	if err != nil {
		err = fmt.Errorf("couldn't open the csv file %v", err)
		return
	}

	return
}

func readProblems(csvfile io.Reader) (problems []problem, err error) {

	reader := csv.NewReader(csvfile)

	records, _ := reader.ReadAll()

	problems = make([]problem, len(records))

	for index, record := range records {
		problems[index] = toProblem(record)
	}

	return
}

func askQuestions(problems []problem, out io.Writer, in io.Reader) results {
	result := results{len(problems), 0}

	reader := bufio.NewReader(in)
	for index, thisProblem := range problems {
		_, _ = fmt.Fprintf(out, "%2d: %s ? ", index+1, thisProblem.question)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if (len(thisProblem.answer) == len(answer)) && (thisProblem.answer == answer) {
			result.right++
		}

		_, _ = fmt.Fprintln(out)
	}

	return result
}

func (r results) show(out io.Writer) {
	_, _ = fmt.Fprintf(out, "Answered %d of %d correctly, %0.0f%% correct\n", r.right, r.asked, float32(r.right)/float32(r.asked)*100.0)
}
