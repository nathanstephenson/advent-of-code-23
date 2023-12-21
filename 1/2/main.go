package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numberStrings = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}

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
	for i, r := range line {
		_, err := strconv.Atoi(string(r))
		if err == nil {
			numberChars = append(numberChars, string(r))
			continue
		}
		charString := string(r)
		for key, value := range numberStrings {
			if i > len(line)-len(key) {
				continue
			}
			if string(key[0]) != charString && key[1] != line[i+1] && key[2] != line[i+2] {
				continue
			}
			length := len(key)
			possibleWord := line[i : i+length]
			if possibleWord != key {
				continue
			}
			numberChars = append(numberChars, fmt.Sprint(value))
		}
	}
	output, err := strconv.Atoi(numberChars[0] + numberChars[len(numberChars)-1])
	if err != nil {
		log.Fatal(err)
	}
	return output
}
