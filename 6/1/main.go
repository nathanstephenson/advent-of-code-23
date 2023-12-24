package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	times := []int{}
	distances := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}
		if timeStr := strings.TrimPrefix(line, "Time: "); timeStr != line {
			for _, v := range strings.Split(timeStr, " ") {
				trimmed := strings.TrimSpace(v)
				if trimmed == "" {
					continue
				}
				times = append(times, aToI(trimmed))
			}
			continue
		}
		if distanceStr := strings.TrimPrefix(line, "Distance: "); distanceStr != line {
			for _, v := range strings.Split(distanceStr, " ") {
				trimmed := strings.TrimSpace(v)
				if trimmed == "" {
					continue
				}
				distances = append(distances, aToI(trimmed))
			}
			continue
		}
	}

	allPossibilities := 0
	for idx, time := range times {
		distance := distances[idx]
		possibleWins := 0
		for chargeTime := 0; chargeTime <= time; chargeTime++ {
			if chargeTime*(time-chargeTime) > distance {
				possibleWins++
			}
		}
		if allPossibilities == 0 {
			allPossibilities = possibleWins
		} else {
			allPossibilities *= possibleWins
		}
	}
	endTime := time.Now().UnixMilli()
	fmt.Printf("Final Output: %d (took %dms)\n", allPossibilities, endTime-startTime)
}

func aToI(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
