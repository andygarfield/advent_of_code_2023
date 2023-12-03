package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type numberLocation struct {
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

	numbers := []numberLocation{}

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
					numberLocation{
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
				numberLocation{
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
		if isPartNumber(input, num) {
			total += toInt(input, num)
		}
	}
	fmt.Println(total)

	// should be 525911
}

func isPartNumber(input [][]rune, num numberLocation) bool {
	for _, indexSet := range findSurroundingIndexes(input, num) {
		checkingChar := input[indexSet[0]][indexSet[1]]
		if checkingChar != '.' && !isDigit(checkingChar) {
			return true
		}
	}

	return false
}

func findSurroundingIndexes(input [][]rune, num numberLocation) [][2]int {
	indices := [][2]int{}

	var startIndex, endIndex int
	if num.startIndex == 0 {
		startIndex = 0
	} else {
		startIndex = num.startIndex - 1
	}

	maxXIndex := len(input[0])
	numEndIndex := num.startIndex + num.len
	if numEndIndex == maxXIndex {
		endIndex = maxXIndex
	} else {
		endIndex = numEndIndex + 1
	}

	maxLineNumber := len(input) - 1
	var startLine, endLine int

	if num.lineIndex == 0 {
		startLine = 0
	} else {
		startLine = num.lineIndex - 1
	}

	if num.lineIndex == maxLineNumber {
		endLine = maxLineNumber
	} else {
		endLine = num.lineIndex + 1
	}

	for i := startLine; i <= endLine; i++ {
		if i == num.lineIndex {
			indices = append(indices, [2]int{i, startIndex}, [2]int{i, endIndex - 1})
		} else {
			for j := startIndex; j < endIndex; j++ {
				indices = append(indices, [2]int{i, j})
			}
		}
	}

	return indices
}

func toInt(input [][]rune, num numberLocation) int {
	i, err := strconv.Atoi(string(input[num.lineIndex][num.startIndex : num.startIndex+num.len]))
	if err != nil {
		log.Panic(err)
	}
	return i
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
