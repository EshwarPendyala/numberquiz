package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file for quiz in the form 'Question, Answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Unable to open file : %s", *csvFilename))
	}
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Unable to read file : %s", *csvFilename))
	}

	problems := parseDataFromLines(lines)
	score := 0
	for i, p := range problems {
		fmt.Printf("Problem No.%d : %s\n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			score++
		}
	}
	fmt.Printf("Your Score is %d out of %d\n", score, len(lines))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func parseDataFromLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}
