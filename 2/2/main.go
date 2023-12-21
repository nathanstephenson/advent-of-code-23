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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameId, power := getGamePower(line)
		fmt.Printf("Game %d power: %v \n", gameId, power)
		sum += power
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Final Output: ", sum)
}

func getGamePower(line string) (int, int) {
	gamesFromId := strings.Split(line, ": ")
	gameIdString := strings.Split(gamesFromId[0], " ")[1]
	gameId, err := strconv.Atoi(gameIdString)
	if err != nil {
		log.Fatal(err)
	}
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	for _, v := range strings.Split(gamesFromId[1], "; ") {
		scores := strings.Split(v, ", ")
		red, green, blue := getColourScores(scores)
		if maxRed < red {
			maxRed = red
		}
		if maxGreen < green {
			maxGreen = green
		}
		if maxBlue < blue {
			maxBlue = blue
		}
	}
	return gameId, maxRed * maxGreen * maxBlue
}

func getColourScores(round []string) (r int, g int, b int) {
	scores := map[string]int{}
	for _, v := range round {
		scoresByColour := strings.Split(v, " ")
		score, err := strconv.Atoi(scoresByColour[0])
		if err != nil {
			log.Fatal(err)
		}
		colour := scoresByColour[1]
		scores[colour] = score
	}
	return scores["red"], scores["green"], scores["blue"]
}
