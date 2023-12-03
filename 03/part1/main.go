package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type number struct {
	lineIndex, startIndex, len int
}

func main() {
	f, err := os.Open("03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := [][]rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}

	if len(input) == 0 {
		log.Fatal("input must contain at least 1 line")
	}

	numbers := []number{}

	currentNumber := ""
	for i := 0; i < len(input); i++ {
		line := input[i]
		for j := 0; j < len(line); j++ {
			thisChar := input[i][j]
			if isDigit(thisChar) {
				currentNumber += string(thisChar)
			} else if len(currentNumber) > 0 {
				numbers = append(
					numbers,
					number{
						lineIndex:  i,
						startIndex: j - len(currentNumber),
						len:        len(currentNumber),
					},
				)
				currentNumber = ""
			}
		}
		if len(currentNumber) > 0 {
			numbers = append(
				numbers,
				number{
					lineIndex:  i,
					startIndex: len(line) - len(currentNumber),
					len:        len(currentNumber),
				},
			)
			currentNumber = ""
		}
	}

	total := 0
	for _, num := range numbers {
		if num.isPartNumber(input) {
			total += num.int(input)
		}
	}
	fmt.Println(total)

	// should be 525911
}

func (n number) isPartNumber(input [][]rune) bool {
	for _, indexSet := range n.findSurroundingIndexes(input) {
		checkingChar := input[indexSet[0]][indexSet[1]]
		if checkingChar != '.' && !isDigit(checkingChar) {
			return true
		}
	}

	return false
}

func (n number) findSurroundingIndexes(input [][]rune) [][2]int {
	indices := [][2]int{}

	var startIndex, endIndex int
	if n.startIndex == 0 {
		startIndex = 0
	} else {
		startIndex = n.startIndex - 1
	}

	maxXIndex := len(input[0])
	numEndIndex := n.startIndex + n.len
	if numEndIndex == maxXIndex {
		endIndex = maxXIndex
	} else {
		endIndex = numEndIndex + 1
	}

	maxLineNumber := len(input) - 1
	var startLine, endLine int

	if n.lineIndex == 0 {
		startLine = 0
	} else {
		startLine = n.lineIndex - 1
	}

	if n.lineIndex == maxLineNumber {
		endLine = maxLineNumber
	} else {
		endLine = n.lineIndex + 1
	}

	for i := startLine; i <= endLine; i++ {
		if i == n.lineIndex {
			indices = append(indices, [2]int{i, startIndex}, [2]int{i, endIndex - 1})
		} else {
			for j := startIndex; j < endIndex; j++ {
				indices = append(indices, [2]int{i, j})
			}
		}
	}

	return indices
}

func (n number) int(input [][]rune) int {
	i, err := strconv.Atoi(string(input[n.lineIndex][n.startIndex : n.startIndex+n.len]))
	if err != nil {
		log.Panic(err)
	}
	return i
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
