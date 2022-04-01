package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func endGame(score int) {
	fmt.Printf("You scored %d points\n", score)
	os.Exit(0)
}
func main() {
	limit := flag.Int("limit", 20, "number of seconds the game will run for")
	csvPath := flag.String("csv", "./questions.csv", "CSV file with questions and answers")
	flag.Parse()
	lim := *limit
	cPath := *csvPath

	fmt.Printf("Limit %d \t filePath : %s\n", lim, cPath)

	file, err := os.Open(cPath)
	if err != nil {
		log.Fatalf("Error opening file %s", cPath)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatalf("Error parsing file %s", cPath)
	}
	go func() {
		time.Sleep(time.Duration(lim) * time.Second)
		fmt.Println("Time up, Bye bye")
		os.Exit(0)
	}()

	var ans string
	score := 0
	input := bufio.NewScanner(os.Stdin)
	for idx, fields := range data {
		fmt.Printf("Q%d. %s\n", idx+1, fields[0])
		input.Scan()
		ans = input.Text()
		if ans == fields[1] {
			fmt.Printf("Correct Answer \u2705 \n")
			score++
		} else {
			fmt.Printf("Incorrect Answer \u274C \n")
		}
	}
	endGame(score)
}
