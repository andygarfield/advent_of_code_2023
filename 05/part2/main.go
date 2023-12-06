package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type itemRange struct {
	destination, source, len int
}

type seedRange struct {
	start, len int
}

func main() {
	f, err := os.Open("05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	seeds, err := parseSeeds(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	mapRanges, err := parseMaps(scanner)
	if err != nil {
		log.Fatal(err)
	}

	smallestLoc := findSmallestLocation(seeds, mapRanges)

	fmt.Println(smallestLoc)
}

func findSmallestLocation(seedRanges []seedRange, mapRanges [][]itemRange) int {
	smallestLoc := 0
	for i, sr := range seedRanges {
		location := findLowestForSeedRange(sr, mapRanges)
		if i == 0 {
			smallestLoc = location
			continue
		}

		if location < smallestLoc {
			smallestLoc = location
		}
	}

	return smallestLoc
}

func findLowestForSeedRange(sr seedRange, mapRanges [][]itemRange) int {
	lowest := 0
	for seed := sr.start; seed < sr.start+sr.len; seed++ {
		currentNumber := seed
		for _, mapList := range mapRanges {
			for _, thisRange := range mapList {
				if thisRange.sourceInRange(currentNumber) {
					currentNumber -= thisRange.source - thisRange.destination
					break
				}
			}
		}

		if seed == sr.start {
			lowest = currentNumber
		}

		if currentNumber < lowest {
			lowest = currentNumber
		}
	}

	return lowest
}

func parseSeeds(line string) ([]seedRange, error) {
	ranges := []seedRange{}

	seedStrings := strings.Split(line[strings.Index(line, " ")+1:], " ")
	for i := 0; i < len(seedStrings); i += 2 {
		start, err := strconv.Atoi(seedStrings[i])
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(seedStrings[i+1])
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, seedRange{start: start, len: length})
	}

	return ranges, nil
}

func parseMaps(scanner *bufio.Scanner) ([][]itemRange, error) {
	mapRanges := [][]itemRange{}

	itemRangeSlice := []itemRange{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// new map starting
		if !isDigit(line[0]) {
			mapRanges = append(mapRanges, itemRangeSlice)
			itemRangeSlice = []itemRange{}
			continue
		}

		itemRangeInts := []int{}
		for _, s := range strings.Split(line, " ") {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			itemRangeInts = append(itemRangeInts, num)
		}

		itemRangeSlice = append(
			itemRangeSlice,
			itemRange{
				destination: itemRangeInts[0],
				source:      itemRangeInts[1],
				len:         itemRangeInts[2],
			},
		)
	}
	mapRanges = append(mapRanges, itemRangeSlice)

	return mapRanges, nil
}

func (i itemRange) sourceInRange(num int) bool {
	return num >= i.source && num < i.source+i.len
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
