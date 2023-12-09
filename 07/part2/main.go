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

type cardType int

func (c cardType) string() string {
	return []string{
		"highCard",
		"onePair",
		"twoPair",
		"threeOfAKind",
		"fullHouse",
		"fourOfAKind",
		"fiveOfAKind",
	}[c]
}

var ranking = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

const (
	highCard cardType = iota
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
		return sortByRanking(a, b)
	}
}

func findType(h hand) cardType {
	labelCount := map[rune]int{}
	for _, r := range h.contents {
		if _, ok := labelCount[r]; !ok {
			labelCount[r] = 0
		}
		labelCount[r]++
	}
	jokerCount := labelCount['J']

	type labelInfo struct {
		label rune
		count int
	}
	counts := []labelInfo{}
	for label, count := range labelCount {
		counts = append(counts, labelInfo{label: label, count: count})
	}
	slices.SortFunc(counts, func(a, b labelInfo) int {
		if a.count > b.count {
			return 1
		} else if a.count < b.count {
			return -1
		} else {
			return 0
		}
	})

	hasThreeOfAKind := false
	pairCount := 0
	for i := len(counts) - 1; i >= 0; i-- {
		count := counts[i]
		tempJokerCount := jokerCount
		if count.label == 'J' {
			if len(counts) > 1 {
				continue
			}
			jokerCount = 0
		}

		if jokerCount+count.count == 5 {
			return fiveOfAKind
		} else if jokerCount+count.count == 4 {
			return fourOfAKind
		} else if jokerCount+count.count == 3 {
			hasThreeOfAKind = true
		} else if jokerCount+count.count == 2 {
			pairCount += 1
		}

		if count.label == 'J' {
			jokerCount = tempJokerCount
			continue
		}

		if jokerCount > 0 {
			jokerCount = 0
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

func sortByRanking(a, b hand) int {
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
