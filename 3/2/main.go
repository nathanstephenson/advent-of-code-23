package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	symbolsByLine := map[int][]int{}
	for i, v := range lines {
		symbolsByLine[i] = getSymbolIndexes(v)
	}
	partsByLine := map[int]map[int]int{}
	for i, line := range lines {
		partsByLine[i] = getValidPartNumbers(i, line, symbolsByLine)
	}

	sum := getGearRatios(symbolsByLine, partsByLine)
	log.Println("Final Output: ", sum)
}

func getGearRatios(symbolsByLine map[int][]int, partsByLine map[int]map[int]int) int {
	sum := 0
	for lineNumber, symbols := range symbolsByLine {
		for _, symbolIdx := range symbols {
			ratio := getGearRatio(lineNumber, symbolIdx, partsByLine)
			if ratio != -1 {
				sum += ratio
			}
		}
	}
	return sum
}

func getGearRatio(lineNumber int, lineIdx int, partsByLine map[int]map[int]int) int {
	adjacentParts := []int{}
	lines := []map[int]int{partsByLine[lineNumber-1], partsByLine[lineNumber], partsByLine[lineNumber+1]}
	for _, parts := range lines {
		endIdx := lineIdx + 1
		for partStart, part := range parts {
			partLen := len(strconv.Itoa(part))
			startIdx := lineIdx - partLen
			fmt.Printf("%d:%d - %d (%d) - validStart=%d validEnd=%d partStart=%d\n", lineNumber, lineIdx, part, partLen, startIdx, endIdx, partStart)
			if startIdx <= partStart && partStart <= endIdx {
				adjacentParts = append(adjacentParts, part)
			}
		}
	}
	fmt.Printf("Part %d:%d - %v\n", lineNumber, lineIdx, adjacentParts)
	if len(adjacentParts) != 2 {
		return -1
	}
	return adjacentParts[0] * adjacentParts[1]
}

func getSymbolIndexes(line string) []int {
	symbols := []int{}
	for i, v := range line {
		matched, err := regexp.MatchString("[*]", string(v))
		if err == nil && matched {
			symbols = append(symbols, i)
		}
	}
	return symbols
}

func getValidPartNumbers(index int, line string, symbolsByLine map[int][]int) map[int]int {
	validParts := map[int]int{}
	partialPart := ""
	partIsValid := false
	for i, v := range line {
		isNumber, err := regexp.MatchString("[0-9]", string(v))
		if err != nil {
			log.Fatal(err)
		}
		if isNumber {
			partialPart += string(v)
			partIsValid = partIsValid || isSymbolAdjacent(index, i, line, symbolsByLine)
		}
		if !isNumber || i == len(line)-1 {
			if partialPart != "" && partIsValid {
				partLen := len(partialPart)
				partNumber, err := strconv.Atoi(partialPart)
				if err != nil {
					log.Fatal(err)
				}
				startIdx := i - partLen
				if i == len(line)-1 && isNumber {
					startIdx = i - (partLen - 1)
				}
				fmt.Printf("Part %d at idx %d\n", partNumber, startIdx)
				validParts[startIdx] = partNumber
			}
			partialPart = ""
			partIsValid = false
		}
	}
	return validParts
}

func isSymbolAdjacent(lineNumber int, lineIndex int, line string, symbolsByLine map[int][]int) bool {
	prevLineSymbols := symbolsByLine[lineNumber-1]
	if prevLineSymbols != nil && lineIndex != 0 && slices.Contains(prevLineSymbols, lineIndex-1) {
		//fmt.Println("Up left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if prevLineSymbols != nil && lineIndex != 0 && slices.Contains(prevLineSymbols, lineIndex) {
		//fmt.Println("Up from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if prevLineSymbols != nil && lineIndex != len(line) && slices.Contains(prevLineSymbols, lineIndex+1) {
		//fmt.Println("Up right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}

	currLineSymbols := symbolsByLine[lineNumber]
	if currLineSymbols != nil && lineIndex != 0 && slices.Contains(currLineSymbols, lineIndex-1) {
		//fmt.Println("Left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if currLineSymbols != nil && lineIndex != len(line) && slices.Contains(currLineSymbols, lineIndex+1) {
		//fmt.Println("Right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}

	nextLineSymbols := symbolsByLine[lineNumber+1]
	if nextLineSymbols != nil && lineIndex != 0 && slices.Contains(nextLineSymbols, lineIndex-1) {
		//fmt.Println("Down left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if nextLineSymbols != nil && slices.Contains(nextLineSymbols, lineIndex) {
		//fmt.Println("Down from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if nextLineSymbols != nil && lineIndex != len(line) && slices.Contains(nextLineSymbols, lineIndex+1) {
		//fmt.Println("Down right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	return false
}
