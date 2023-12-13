package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type pos struct{ x, y int }

func main() {
	f, err := os.Open("11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	galaxies, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	result := solve(galaxies)
	fmt.Println(result)
}

func solve(galaxies []pos) int {
	type indexPair struct{ galaxy1, galaxy2 int }
	distances := map[indexPair]int{}
	for i, galaxy1 := range galaxies {
		for j, galaxy2 := range galaxies {
			if _, ok := distances[indexPair{galaxy1: j, galaxy2: i}]; ok {
				continue
			}

			distances[indexPair{galaxy1: i, galaxy2: j}] = int(
				math.Abs(
					float64(galaxy2.x-galaxy1.x),
				) + math.Abs(
					float64(galaxy2.y-galaxy1.y),
				),
			)
		}
	}

	total := 0
	for _, distance := range distances {
		total += distance
	}

	return total
}

func parseInput(r io.ReadSeeker) ([]pos, error) {
	scanner := bufio.NewScanner(r)

	blankRows := 0
	blankColMap := map[int]struct{}{}
	galaxies := []pos{}
	for y := 0; scanner.Scan(); y++ {
		rowIsBlank := true
		row := scanner.Bytes()

		for x, value := range row {
			if y == 0 {
				blankColMap[x] = struct{}{}
			}
			if value != '.' {
				rowIsBlank = false
				delete(blankColMap, x)
				galaxies = append(galaxies, pos{x: x, y: y + blankRows})
			}
		}

		if rowIsBlank {
			blankRows++
		}
	}

	for galaxyIndex, galaxy := range galaxies {
		for blankCol := range blankColMap {
			if blankCol < galaxy.x {
				galaxies[galaxyIndex].x++
			}
		}
	}

	return galaxies, nil
}
