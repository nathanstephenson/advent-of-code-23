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
	scores := map[int]int{}
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		scores[lineNum] = getScoreForCard(line)
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Scores: %v\n", scores)
	sum := getTotalCardCount(scores)
	log.Println("Final Output: ", sum)
}

func getTotalCardCount(scores map[int]int) int {
	cards := map[int]int{}
	sum := 0
	for i := 1; i <= len(scores); i++ {
		v := scores[i]
		cards[i]++
		for j := 1; j <= v; j++ {
			cards[i+j] += (1 * cards[i])
		}
		fmt.Printf("Card: %d (%d) - %d\n", i, scores[i], cards[i])
		sum += cards[i]
	}
	return sum
}

func getScoreForCard(line string) int {
	_, numbers := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
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
	score := 0
	for _, v := range picked {
		if slices.Contains(winning, v) {
			score++
		}
	}
	return score
}
