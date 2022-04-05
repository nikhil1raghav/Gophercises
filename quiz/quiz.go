package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"
)

type problem struct {
	que string
	ans string
}

func main() {
	csvFile := flag.String("csv", "questions.csv", "csv file that has quiz data")
	timeLimit := flag.Int("limit", 30, "time in seconds that quiz will run for")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Println("Error opening csv file : ", *csvFile)
		os.Exit(1)
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error in parsing csv file")
		os.Exit(1)
	}
	problems := parseLines(lines)
	score := 0
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	timer := time.NewTimer(time.Duration(*timeLimit * int(time.Second)))
	for id, prob := range problems {
		fmt.Printf("Problem #%d: %s  \n", id+1, prob.que)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scan(&answer)
			answerChan <- answer

		}()
		select {
		case <-timer.C:
			endGame(score, len(problems))
		case <-sigChan:
			endGame(score, len(problems))
		case answer := <-answerChan:
			if answer == prob.ans {
				score++
				fmt.Println("Correct \u2705")
			} else {
				fmt.Println("Incorrect \u274c")
			}
		}
	}
}
func endGame(score, total int) {
	fmt.Printf("You scored %d out of %d\n", score, total)
	os.Exit(0)
}

func parseLines(lines [][]string) []problem {
	var problems []problem
	for _, line := range lines {
		if len(line) < 2 {
			fmt.Println("Less elements to form a complete problem, skipping...")
			continue
		}
		problems = append(problems, problem{que: line[0], ans: strings.TrimSpace(line[1])})
	}
	return problems

}
