package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allNumbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		numberFromLine := getNumbersFromLine(line)
		fmt.Printf("%d from line %s \n", numberFromLine, line)
		allNumbers = append(allNumbers, numberFromLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range allNumbers {
		sum += v
	}
	log.Println("Final Output: ", sum)
}

func getNumbersFromLine(line string) int {
	numberChars := []string{}
	for _, r := range line {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			numberChars = append(numberChars, string(r))
		}
	}
	output, err := strconv.Atoi(numberChars[0] + numberChars[len(numberChars)-1])
	if err != nil {
		log.Fatal(err)
	}
	return output
}
