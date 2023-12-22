package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapping := map[string][][3]int{}
	lastMap := ""
	seeds := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}
		if seedsStr := strings.TrimPrefix(line, "seeds: "); seedsStr != line {
			seeds = getSeeds(seedsStr)
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

	locations := []int{}
	for _, seed := range seeds {
		location := seed
		for _, mapName := range maps {
			for _, v := range mapping[mapName] {
				dest, src, r := v[0], v[1], v[2]-1
				if location < src || src+r < location {
					continue
				}
				offset := location - src
				location = dest + offset
				//fmt.Printf("Seed %d: %v, dest %d (offset %d)\n", seed, mapName, location, offset)
				break
			}
		}
		locations = append(locations, location)
	}

	log.Println("Final Output: ", slices.Min(locations))
}

func getSeeds(seedsStr string) []int {
	seeds := []int{}
	for _, v := range strings.Split(seedsStr, " ") {
		if seedStr := strings.TrimSpace(v); seedsStr != "" {
			seed, err := strconv.Atoi(seedStr)
			if err != nil {
				log.Fatal(err)
			}
			seeds = append(seeds, seed)
		}
	}
	return seeds
}
