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
	t := 0
	d := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}
		if timesStr := strings.TrimPrefix(line, "Time: "); timesStr != line {
			timeStr := ""
			for _, v := range strings.Split(timesStr, " ") {
				trimmed := strings.TrimSpace(v)
				if trimmed == "" {
					continue
				}
				timeStr += trimmed
			}
			t = aToI(timeStr)
			continue
		}
		if distancesStr := strings.TrimPrefix(line, "Distance: "); distancesStr != line {
			distanceStr := ""
			for _, v := range strings.Split(distancesStr, " ") {
				trimmed := strings.TrimSpace(v)
				if trimmed == "" {
					continue
				}
				distanceStr += trimmed
			}
			d = aToI(distanceStr)
			continue
		}
	}

	allPossibilities := 0
	possibleWins := 0
	for chargeTime := 0; chargeTime <= t; chargeTime++ {
		if chargeTime*(t-chargeTime) > d {
			possibleWins++
		}
	}
	if allPossibilities == 0 {
		allPossibilities = possibleWins
	} else {
		allPossibilities *= possibleWins
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
