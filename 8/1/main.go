package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ScanBuf(scanner *bufio.Scanner, handleLine func(string)) {
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line)
	}
}

type Pair[T, U any] struct {
	Left  T
	Right U
}

func main() {
	startTime := time.Now().UnixMilli()
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	steps := ""
	inputs := map[string]Pair[string, string]{}
	firstLine := true
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		if firstLine {
			steps = line
			firstLine = false
			return
		}
		splitLine := strings.Split(line, " = ")
		idx, leftRight := splitLine[0], strings.Split(strings.TrimSuffix(strings.TrimPrefix(splitLine[1], "("), ")"), ", ")
		left, right := leftRight[0], leftRight[1]
		inputs[idx] = Pair[string, string]{left, right}
	}
	ScanBuf(scanner, handleLine)

	stepCount := getSteps(steps, inputs)
	endTime := time.Now().UnixMilli()
	fmt.Printf("Final Output: %d (took %dms)\n", stepCount, endTime-startTime)
}

func getSteps(steps string, inputs map[string]Pair[string, string]) int {
	start := "AAA"
	dest := "ZZZ"
	count := 0
	current := start
	for {
		for _, s := range steps {
			if current == dest {
				return count
			}
			count++
			if s == 'L' {
				current = inputs[current].Left
			}
			if s == 'R' {
				current = inputs[current].Right
			}
		}
	}
}
