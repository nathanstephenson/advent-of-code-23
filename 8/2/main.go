package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func getSteps(steps string, inputs map[string]Pair[string, string]) int64 {
	isDest := func(s string) bool {
		return strings.HasSuffix(s, "Z")
	}
	startingNodes := getStartingNodes(inputs)
	routeLengths := []Pair[string, int]{}
	for _, v := range startingNodes {
		count := 0
		current := v
		routeLengths = append(routeLengths, Pair[string, int]{v, getStepsForRoute(steps, current, count, inputs, isDest)})
	}
	fmt.Printf("Route Lengths: %v\n", routeLengths)
	return lcm(routeLengths)
}

func lcm(routeLengths []Pair[string, int]) int64 {
	primeFactors := []int64{}
	for _, v := range routeLengths {
		factors := getPrimeFactors(int64(v.Right))
		primeFactors = append(primeFactors, factors...)
	}
	factors := removeDuplicates(primeFactors)
	var lcm int64 = 1
	for _, v := range factors {
		lcm *= v
	}
	return lcm
}

func removeDuplicates(slice []int64) []int64 {
	sliceMap := map[int64]int{}
	for _, v := range slice {
		sliceMap[v] = 1
	}
	set := make([]int64, len(sliceMap))
	i := 0
	for k := range sliceMap {
		set[i] = k
		i++
	}
	return set
}

func getPrimeFactors(num int64) []int64 {
	primeFactors := []int64{}
	for i := int64(2); i <= num; i++ {
		if num%i == 0 {
			primeFactors = append(primeFactors, i)
			num /= i
			i = 1
		}
		if doFactorsProd(primeFactors, num) {
			return primeFactors
		}
	}
	return primeFactors
}

func areAllFactorsPrime(factors []int64) bool {
	for _, v := range factors {
		if len(getPrimeFactors(v)) > 1 {
			return false
		}
	}
	return true
}

func doFactorsProd(primeFactors []int64, target int64) bool {
	prod := int64(1)
	for _, v := range primeFactors {
		prod *= v
	}
	return prod == target
}

func allMatch(matches []bool, len int) bool {
	expected := []bool{}
	for i := 0; i < len; i++ {
		expected = append(expected, true)
	}
	fmt.Println(matches, expected)
	return slices.Equal(matches, expected)
}

func getStepsForRoute(steps string, current string, count int, inputs map[string]Pair[string, string], isDest func(string) bool) int {
	for {
		for _, s := range steps {
			if isDest(current) {
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

func getStartingNodes(inputs map[string]Pair[string, string]) []string {
	startingNodes := []string{}
	for k := range inputs {
		if strings.HasSuffix(k, "A") {
			startingNodes = append(startingNodes, k)
		}
	}
	return startingNodes
}
