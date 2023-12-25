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

func scanBuf(scanner *bufio.Scanner, handleLine func(string)) {
	for scanner.Scan() {
		line := scanner.Text()
		handleLine(line)
	}
}

func AtoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func runeAt(s string, i int) rune {
	for idx, v := range s {
		if idx == i {
			return v
		}
	}
	return 0
}

type Pair[T, U any] struct {
	First  T
	Second U
}

const FIVE_OF_A_KIND = 6
const FOUR_OF_A_KIND = 5
const FULL_HOUSE = 4
const THREE_OF_A_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

var CARDS = map[rune]int{'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8, '9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0}

func main() {
	startTime := time.Now().UnixMilli()
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	handBidPairs := [](Pair[string, int]){}
	handleLine := func(line string) {
		if strings.Trim(line, " ") == "" {
			return
		}
		splitLine := strings.Split(line, " ")
		pair := Pair[string, int]{splitLine[0], AtoI(splitLine[1])}
		handBidPairs = append(handBidPairs, pair)
	}
	scanBuf(scanner, handleLine)

	winnings := calculateWinnings(handBidPairs)

	endTime := time.Now().UnixMilli()
	fmt.Printf("Final Output: %d (took %dms)\n", winnings, endTime-startTime)
}

func calculateWinnings(handBidPairs []Pair[string, int]) int {
	scores := map[string]int{}
	for _, p := range handBidPairs {
		cardCounts := organiseHand(p.First)
		score := scoreHand(cardCounts)
		scores[p.First] = score
	}
	sortedHand := sortHandsByScore(handBidPairs, scores)
	sum := 0
	for i, p := range sortedHand {
		sum += (i + 1) * p.Second
	}
	return sum
}

func sortHandsByScore(handBidPairs []Pair[string, int], scores map[string]int) []Pair[string, int] {
	sortedHand := slices.Clone(handBidPairs)
	slices.SortFunc(sortedHand, func(i, j Pair[string, int]) int {
		if scores[i.First] > scores[j.First] {
			return 1
		}
		if scores[i.First] < scores[j.First] {
			return -1
		}
		for idx, v := range i.First {
			jRune := runeAt(j.First, idx)
			if CARDS[v] > CARDS[jRune] {
				return 1
			}
			if CARDS[v] < CARDS[jRune] {
				return -1
			}
		}
		return 0
	})
	return sortedHand
}

func scoreHand(sortedCards map[rune]int) int {
	if len(sortedCards) == 1 {
		return FIVE_OF_A_KIND
	}
	if len(sortedCards) == 2 {
		for _, v := range sortedCards {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}
	if len(sortedCards) == 3 {
		for _, v := range sortedCards {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	}
	if len(sortedCards) == 4 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func organiseHand(hand string) map[rune]int {
	cardCounts := map[rune]int{}
	for _, v := range hand {
		val := cardCounts[v]
		cardCounts[v] = val + 1
	}
	return cardCounts
}
