package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	history, err := parseInput(f)
	if err != nil {
		log.Fatal(err)
	}

	result := solve(history)
	fmt.Println(result)
}

func solve(history [][]int) int {
	sum := 0
	for _, single := range history {
		sum += solveSingle(single)
	}
	return sum
}

func solveSingle(history []int) int {
	reductions := [][]int{history}
	lastReduction := history

	for !isAllZeros(lastReduction) {
		reduction := make([]int, len(lastReduction)-1)
		for i := 0; i < len(lastReduction)-1; i++ {
			reduction[i] = lastReduction[i+1] - lastReduction[i]
		}
		reductions = append(reductions, reduction)
		lastReduction = reductions[len(reductions)-1]
	}

	lastReduction = append(lastReduction, 0)

	for i := len(reductions) - 1; i > 0; i-- {
		reductions[i-1] = append(
			reductions[i-1],
			reductions[i-1][len(reductions[i-1])-1]+reductions[i][len(reductions[i])-1],
		)
	}

	return reductions[0][len(reductions[0])-1]
}

func isAllZeros(ints []int) bool {
	for _, i := range ints {
		if i != 0 {
			return false
		}
	}
	return true
}

func parseInput(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	out := [][]int{}
	for scanner.Scan() {
		ints := []int{}
		for _, intStr := range strings.Split(scanner.Text(), " ") {
			i, err := strconv.Atoi(intStr)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}
		out = append(out, ints)
	}

	return out, nil
}
