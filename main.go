package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Provide a csv file of format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("We were not able to open the file: %s", *fileName)
	}
	defer os.Exit(1)

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Exit. Failed to parse the csv file.")
	}

	problems := parseLines(lines)
	var count int
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			count++
		}
	}
	fmt.Printf("Your final score is: %d\n", count)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
