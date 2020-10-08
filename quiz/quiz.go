package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

// ask asks a question and returns an updated score depending on the answer.
func ask(s chan int, question question, score int) {
	fmt.Println(question.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, question.a) == 0 {
		score++
		s <- score
		fmt.Println("Correct!")
	} else {
		s <- score
		fmt.Println("Incorrect :-(")
	}
}

func main() {
	timer := time.NewTimer(5 * time.Second)
	score := 0
	s := make(chan int)
	qs := questions()
	Loop:
		for _, q := range qs {
			go ask(s,q,score)
			select {
				case x := <- s:
					score = x
				case <- timer.C:
					break Loop
			}
		}
	fmt.Println("Final score", score)
}
