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
	symbolsByLine := map[int][]int{}
	for i, v := range lines {
		symbolsByLine[i] = getSymbolIndexes(v)
	}
	sum := 0
	for i, line := range lines {
		validParts := getValidPartNumbers(i, line, symbolsByLine)
		fmt.Printf("Line %d has valid parts: %v\n", i, validParts)
		for _, part := range validParts {
			sum += part
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Final Output: ", sum)
}

func getSymbolIndexes(line string) []int {
	symbols := []int{}
	for i, v := range line {
		matched, err := regexp.MatchString("[^a-zA-Z0-9\\s.]", string(v))
		if err == nil && matched {
			symbols = append(symbols, i)
		}
	}
	return symbols
}

func getValidPartNumbers(index int, line string, symbolsByLine map[int][]int) []int {
	validParts := []int{}
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
				partNumber, err := strconv.Atoi(partialPart)
				if err != nil {
					log.Fatal(err)
				}
				validParts = append(validParts, partNumber)
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
		fmt.Println("Up left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if prevLineSymbols != nil && lineIndex != 0 && slices.Contains(prevLineSymbols, lineIndex) {
		fmt.Println("Up from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if prevLineSymbols != nil && lineIndex != len(line) && slices.Contains(prevLineSymbols, lineIndex+1) {
		fmt.Println("Up right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}

	currLineSymbols := symbolsByLine[lineNumber]
	if currLineSymbols != nil && lineIndex != 0 && slices.Contains(currLineSymbols, lineIndex-1) {
		fmt.Println("Left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if currLineSymbols != nil && lineIndex != len(line) && slices.Contains(currLineSymbols, lineIndex+1) {
		fmt.Println("Right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}

	nextLineSymbols := symbolsByLine[lineNumber+1]
	if nextLineSymbols != nil && lineIndex != 0 && slices.Contains(nextLineSymbols, lineIndex-1) {
		fmt.Println("Down left from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if nextLineSymbols != nil && slices.Contains(nextLineSymbols, lineIndex) {
		fmt.Println("Down from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	if nextLineSymbols != nil && lineIndex != len(line) && slices.Contains(nextLineSymbols, lineIndex+1) {
		fmt.Println("Down right from ", lineNumber, ":", lineIndex, " has a symbol")
		return true
	}
	return false
}
