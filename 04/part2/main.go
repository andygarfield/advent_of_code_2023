package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	matching int
	count    int
}

func main() {
	f, err := os.Open("04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	allCards := []card{}
	cardNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		cardNum += 1

		numbersString := line[strings.Index(line, ":")+1:]

		numberSplit := strings.Split(numbersString, "|")
		winningNumberList := numberSplit[0]

		winningNumbers := map[int]struct{}{}
		for _, numString := range strings.Split(winningNumberList, " ") {
			if len(numString) == 0 {
				continue
			}

			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}

			winningNumbers[num] = struct{}{}
		}

		thisCard := card{matching: 0, count: 1}
		for _, numString := range strings.Split(numberSplit[1], " ") {
			if len(numString) == 0 {
				continue
			}

			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}

			if _, exists := winningNumbers[num]; exists {
				thisCard.matching += 1
			}
		}
		allCards = append(allCards, thisCard)
	}

	for cardIndex, card := range allCards {
		for countIteration := 0; countIteration < card.count; countIteration++ {
			maxCopyIndex := cardIndex + card.matching
			for toCopyIndex := cardIndex + 1; toCopyIndex < maxCopyIndex+1 && toCopyIndex < len(allCards); toCopyIndex++ {
				allCards[toCopyIndex].count += 1
			}
		}
	}

	totalCards := 0
	for _, card := range allCards {
		totalCards += card.count
	}

	fmt.Println(totalCards)
}
