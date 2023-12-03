package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type schematic struct {
	input                        [][]rune
	currentLine, currentLineChar uint
}

func main() {
	f, err := os.Open("03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	schematic, err := NewSchematic(f)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for schematic.hasNext() {
		total += schematic.next()
	}
	fmt.Println(total)
	// should be 4473
}

func NewSchematic(r io.Reader) (*schematic, error) {
	input := [][]rune{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}

	if len(input) == 0 {
		return nil, errors.New("input must contain at least 1 line")
	}

	return &schematic{input: input}, nil
}

func (s *schematic) hasNext() bool {
	return s.lineInBounds() && s.charInBounds()
}

func (s *schematic) charInBounds() bool {
	if len(s.input[0]) == int(s.currentLineChar) {
		return false
	} else {
		return true
	}
}

func (s *schematic) lineInBounds() bool {
	if len(s.input) == int(s.currentLine) {
		return false
	} else {
		return true
	}
}

func (s *schematic) next() int {
	currentNumber := ""
	getPartNumber := func() (int, bool) {
		if len(currentNumber) > 0 {
			isPartNumber := s.isPartNumber(
				s.currentLine,
				s.currentLineChar-uint(len(currentNumber)),
				s.currentLineChar,
			)
			if isPartNumber {
				num, err := strconv.Atoi(currentNumber)
				if err != nil {
					log.Panic(err)
				}
				return num, true
			} else {
				currentNumber = ""
				return 0, false
			}
		}

		return 0, false
	}

	for s.lineInBounds() {
		for s.charInBounds() {
			char := s.input[s.currentLine][s.currentLineChar]
			if isDigit(char) {
				currentNumber += string(char)
			} else {
				partNumber, isPartNumber := getPartNumber()
				if isPartNumber {
					return partNumber
				}
			}

			s.currentLineChar++
		}
		partNumber, isPartNumber := getPartNumber()

		s.currentLineChar = 0
		s.currentLine++

		if isPartNumber {
			return partNumber
		}
	}

	return 0
}

func (s *schematic) isPartNumber(lineNum, startIndex, toIndex uint) bool {
	for _, indexSet := range s.findSurroundingIndexes(lineNum, startIndex, toIndex) {
		checkingChar := s.input[indexSet[0]][indexSet[1]]
		if checkingChar != '.' && !isDigit(checkingChar) {
			return true
		}
	}

	return false
}

func (s *schematic) findSurroundingIndexes(lineNum, startIndex, toIndex uint) [][2]uint {
	indices := [][2]uint{}

	if startIndex != 0 {
		startIndex -= 1
	}

	maxXIndex := uint(len(s.input[0]))
	if toIndex != maxXIndex {
		toIndex += 1
	}

	maxLineNumber := uint(len(s.input) - 1)
	var startLine, endLine uint
	if lineNum == 0 {
		startLine = 0
	} else {
		startLine = lineNum - 1
	}

	if lineNum == maxLineNumber {
		endLine = maxLineNumber
	} else {
		endLine = lineNum + 1
	}

	//return indices
	for i := startLine; i <= endLine; i++ {
		for j := startIndex; j < toIndex; j++ {
			indices = append(indices, [2]uint{i, j})
		}
	}

	return indices
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
