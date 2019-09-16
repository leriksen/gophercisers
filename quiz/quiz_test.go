package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestProblemToStringHappy(t *testing.T) {
	result := toProblem([]string{"1+1", "2"})

	if result.question != "1+1" {
		t.Error(fmt.Sprintf("parse failure: question - expected 1+1, got %s", result.question))
	}
	if result.answer != "2" {
		t.Error(fmt.Sprintf("parse failure: answer - expected 2, got %s", result.answer))
	}
}

func TestProblemToStringWithSpaces(t *testing.T) {
	result := toProblem([]string{" 1+1 ", " 2 "})

	if result.question != "1+1" {
		t.Error(fmt.Sprintf("parse failure: question - expected '1+1', got '%s'", result.question))
	}
	if result.answer != "2" {
		t.Error(fmt.Sprintf("parse failure: answer - expected '2', got '%s'", result.answer))
	}
}

func TestOpenNonExistentFile(t *testing.T) {
	_, err := openFile("nosuchfile")

	if err == nil {
		t.Error(fmt.Sprintf("open failure: expected non-nil, got '%v'", err))

	}
}

func TestOpenExistentFile(t *testing.T) {
	_, err := openFile(os.Args[0])

	if err != nil {
		t.Error(fmt.Sprintf("open failure: expected nil, got '%v'", err))

	}
}

func TestResultShowAllWrong(t *testing.T) {
	r := results{
		asked: 10,
		right: 0,
	}

	var b bytes.Buffer

	r.show(&b)

	expected := "Answered 0 of 10 correctly, 0% correct\n"
	if b.String() != expected {
		t.Error(fmt.Sprintf("show failure: got '%s', expected '%s'", b.String(), expected))
	}
}

func TestResultShowAllRight(t *testing.T) {
	r := results{
		asked: 10,
		right: 10,
	}

	var b bytes.Buffer

	r.show(&b)

	expected := "Answered 10 of 10 correctly, 100% correct\n"
	if b.String() != expected {
		t.Error(fmt.Sprintf("show failure: got '%s', expected '%s'", b.String(), expected))
	}
}

func TestResultShowOneRight(t *testing.T) {
	r := results{
		asked: 12,
		right: 1,
	}

	var b bytes.Buffer

	r.show(&b)

	expected := "Answered 1 of 12 correctly, 8% correct\n"
	if b.String() != expected {
		t.Error(fmt.Sprintf("show failure: got '%s', expected '%s'", b.String(), expected))
	}
}

func TestAskQuestionsRight(t *testing.T) {
	problems := []problem{
		{
			"1+1",
			"2",
		},
	}

	var in = bytes.NewBufferString("2\n")
	out := bytes.Buffer{}

	results := askQuestions(problems, &out, in)

	if results.asked != 1 {
		t.Error(fmt.Sprintf("askQuestions asked: expected 1, got %d", results.asked))
	}

	if results.right != 1 {
		t.Error(fmt.Sprintf("askQuestions right: expected 1, got %d", results.right))
	}
}

func TestAskQuestionsWrong(t *testing.T) {
	problems := []problem{
		{
			"1+1",
			"2",
		},
	}

	var in = bytes.NewBufferString("1\n")
	out := bytes.Buffer{}

	results := askQuestions(problems, &out, in)

	if results.asked != 1 {
		t.Error(fmt.Sprintf("askQuestions asked: expected 1, got %d", results.asked))
	}

	if results.right != 0 {
		t.Error(fmt.Sprintf("askQuestions right: expected 0, got %d", results.right))
	}
}

func TestReadProblemsOK(t *testing.T) {
	csvfile := bytes.NewBufferString("1+1,2\n2+3,5")

	problems, _ := readProblems(csvfile)

	if len(problems) != 2 {
		t.Error(fmt.Sprintf("readProblems length: expected 2, got %d", len(problems)))
	}

	if problems[0].question != "1+1" {
		t.Error(fmt.Sprintf("readProblems result: expected '1+1', got %s", problems[0].question))
	}
	if problems[1].question != "2+3" {
		t.Error(fmt.Sprintf("readProblems result: expected '2+3', got %s", problems[0].question))
	}
	if problems[0].answer != "2" {
		t.Error(fmt.Sprintf("readProblems result: expected '2', got %s", problems[0].answer))
	}
	if problems[1].answer != "5" {
		t.Error(fmt.Sprintf("readProblems result: expected '5', got %s", problems[0].answer))
	}
}
