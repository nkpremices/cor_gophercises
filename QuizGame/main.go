package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main()  {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, readError := r.ReadAll()

	if readError != nil {
		exit("Failed to parse the provided csv file")
	}

	problems := parseLines(lines)

	fmt.Println(problems)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines (lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type Problem struct {
	q string
	a string
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}