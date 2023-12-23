package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var maps = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func main() {
	startTime := time.Now().UnixMilli()
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapping := map[string][][3]int{}
	lastMap := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}
		if seedsStr := strings.TrimPrefix(line, "seeds: "); seedsStr != line {
			mapping["seeds"] = getSeeds(seedsStr)
			continue
		}
		if mapName := strings.TrimSuffix(line, " map:"); mapName != line {
			lastMap = mapName
			continue
		}
		valuesStr := strings.TrimSpace(line)
		strValues := strings.Split(valuesStr, " ")
		src, err := strconv.Atoi(strings.TrimSpace(strValues[0]))
		if err != nil {
			log.Fatal(err)
		}
		dest, err := strconv.Atoi(strings.TrimSpace(strValues[1]))
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(strings.TrimSpace(strValues[2]))
		if err != nil {
			log.Fatal(err)
		}
		_, exists := mapping[lastMap]
		if !exists {
			mapping[lastMap] = [][3]int{}
		}
		mapping[lastMap] = append(mapping[lastMap], [3]int{src, dest, r})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Processed file, looking for seed of lowest output")
	reversedMaps := slices.Clone(maps)
	slices.Reverse(reversedMaps)
	seed := 0
	startLocation := 0
	fmt.Printf("Looking for lowest seed from starting '%v' %d\n", reversedMaps[0], startLocation)
	for start := 0; !isSeedValid(seed, mapping["seeds"]); start++ {
		startLocation = start
		location := start
		for _, mapName := range reversedMaps {
			for _, v := range mapping[mapName] {
				src, dest, r := v[0], v[1], v[2]-1
				if src <= location && location <= src+r {
					offset := location - src
					location = dest + offset
					break
				}
			}
			if mapName == "seed-to-soil" {
				seed = location
				fmt.Printf("Seed %d: %v\n", seed, start)
			}
		}
	}
	endTime := time.Now().UnixMilli()

	log.Printf("Final Output: seed = %d, location = %d (took %d ms)", seed, startLocation, endTime-startTime)
}

func isSeedValid(seed int, seeds [][3]int) bool {
	for _, v := range seeds {
		start, r := v[0], v[2]
		if start <= seed && seed <= start+r {
			return true
		}
	}
	return false
}

func getSeeds(seedsStr string) [][3]int {
	seeds := [][3]int{}
	seedStrings := strings.Split(seedsStr, " ")
	for i := 0; i < len(seedStrings); i += 2 {
		vStr := seedStrings[i]
		rStr := seedStrings[i+1]
		if seedStr := strings.TrimSpace(vStr); seedsStr != "" {
			if seedRange := strings.TrimSpace(rStr); seedRange != "" {
				seed, err := strconv.Atoi(seedStr)
				if err != nil {
					log.Fatal(err)
				}
				r, err := strconv.Atoi(seedRange)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, [3]int{seed, 0, r})
			}
		}
	}
	return seeds
}
