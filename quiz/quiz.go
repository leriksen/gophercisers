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

	problems := readProblems(*filename)

	result := askQuestions(problems, out, in)

	result.show(out)
}

func (p problem) String() string {
	return fmt.Sprintf("%s,%s", p.question, p.answer)
}

func toProblem(row []string) problem {
	return problem{
		row[0],
		row[1],
	}
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error closing file: %v\n", err)
		os.Exit(1)
	}
}

func readProblems(file string) []problem {
	var problems []problem

	csvfile, err := os.Open(file)

	defer closeFile(csvfile)

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	reader := csv.NewReader(csvfile)

	records, err := reader.ReadAll()

	for _, record := range records {
		problems = append(problems, toProblem(record))
	}
	return problems
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
