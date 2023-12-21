package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validGames := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		gameId, isValidGame := determineIfGameIsValid(line)
		fmt.Printf("%d: Valid? %v \n", gameId, isValidGame)
		if isValidGame {
			validGames = append(validGames, gameId)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range validGames {
		sum += v
	}
	log.Println("Final Output: ", sum)
}

var maxScores = map[string]int{"red": 12, "green": 13, "blue": 14}

func determineIfGameIsValid(line string) (int, bool) {
	gamesFromId := strings.Split(line, ": ")
	gameIdString := strings.Split(gamesFromId[0], " ")[1]
	gameId, err := strconv.Atoi(gameIdString)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range strings.Split(gamesFromId[1], "; ") {
		scores := strings.Split(v, ", ")
		if !isRoundValid(scores) {
			return gameId, false
		}
	}
	return gameId, true
}

func isRoundValid(round []string) bool {
	for _, v := range round {
		scores := strings.Split(v, " ")
		score, err := strconv.Atoi(scores[0])
		if err != nil {
			log.Fatal(err)
		}
		colour := scores[1]
		if score > maxScores[colour] {
			return false
		}
	}
	return true
}
