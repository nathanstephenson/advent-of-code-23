package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += getScoreForCard(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Final Output: ", sum)
}

func getScoreForCard(line string) int {
	cardIdx, numbers := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
	winningStr, pickedStr := strings.Split(numbers, "|")[0], strings.Split(numbers, "|")[1]
	winning := []string{}
	for _, v := range strings.SplitAfter(winningStr, " ") {
		if v == " " || strings.Trim(v, " ") == "" {
			continue
		}
		winning = append(winning, strings.Trim(v, " "))
	}
	picked := []string{}
	for _, v := range strings.SplitAfter(pickedStr, " ") {
		if v == " " || strings.Trim(v, " ") == "" {
			continue
		}
		picked = append(picked, strings.Trim(v, " "))
	}
	fmt.Printf("%v - winning %v (%d) - picked %v (%d)\n", cardIdx, winning, len(winning), picked, len(picked))
	score := 0
	for _, v := range picked {
		if slices.Contains(winning, v) {
			if score == 0 {
				fmt.Printf("Found %v", v)
				score = 1
			} else {
				fmt.Printf(", %v", v)
				score *= 2
			}
		}
	}
	fmt.Printf("\nScore - %d\n", score)
	return score
}
