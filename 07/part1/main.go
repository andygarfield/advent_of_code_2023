package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ranking = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	contents []rune
	bid      int
}

func main() {
	f, err := os.Open("07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	hands, err := parseHands(f)
	if err != nil {
		log.Fatal(err)
	}

	answer := solve(hands)
	fmt.Println(answer)
}

func solve(hands []hand) int {
	slices.SortFunc(hands, sorter)

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}

	return total
}

func sorter(a, b hand) int {
	aType := findType(a)
	bType := findType(b)

	if aType > bType {
		return 1
	} else if aType < bType {
		return -1
	} else {
		return findHigher(a, b)
	}
}

func findType(h hand) int {
	labelCount := map[rune]int{}
	for _, r := range h.contents {
		if _, ok := labelCount[r]; !ok {
			labelCount[r] = 0
		}
		labelCount[r]++
	}
	hasThreeOfAKind := false
	pairCount := 0
	for _, count := range labelCount {
		if count == 5 {
			return fiveOfAKind
		}
		if count == 4 {
			return fourOfAKind
		}
		if count == 3 {
			hasThreeOfAKind = true
		}
		if count == 2 {
			pairCount += 1
		}
	}

	if hasThreeOfAKind && pairCount == 1 {
		return fullHouse
	} else if hasThreeOfAKind {
		return threeOfAKind
	} else if pairCount == 2 {
		return twoPair
	} else if pairCount == 1 {
		return onePair
	} else {
		return highCard
	}
}

func findHigher(a, b hand) int {
	for i, aLabel := range a.contents {
		bLabel := b.contents[i]
		if ranking[aLabel] > ranking[bLabel] {
			return 1
		} else if ranking[aLabel] < ranking[bLabel] {
			return -1
		}
	}

	return 0
}

func parseHands(r io.Reader) ([]hand, error) {
	hands := []hand{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		contents := []rune(split[0])
		bidString := split[1]
		bid, err := strconv.Atoi(bidString)
		if err != nil {
			return nil, err
		}

		hands = append(hands, hand{contents: contents, bid: bid})
	}

	return hands, nil
}
