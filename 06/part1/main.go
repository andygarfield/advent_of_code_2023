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
	f, err := os.Open("06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	times, records, err := parseFile(f)
	if err != nil {
		log.Fatal(err)
	}

	solution := solve(times, records)

	fmt.Println(solution)
}

func solve(times []int, records []int) int {
	var answer int
	for i, time := range times {
		record := records[i]

		timesOver := 0
		for timeHeld := 0; timeHeld <= time; timeHeld++ {
			totalTime := (time - timeHeld) * timeHeld
			if totalTime > record {
				timesOver++
			}
		}
		if i == 0 {
			answer = timesOver
		} else {
			answer *= timesOver
		}
	}

	return answer
}

func parseFile(r io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	timeLine := scanner.Text()[11:]

	scanner.Scan()
	recordsLine := scanner.Text()[11:]

	times := []int{}
	for _, numStr := range strings.Split(timeLine, " ") {
		if len(numStr) == 0 {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, nil, err
		}

		times = append(times, num)
	}

	records := []int{}
	for _, numStr := range strings.Split(recordsLine, " ") {
		if len(numStr) == 0 {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, nil, err
		}

		records = append(records, num)
	}

	return times, records, nil
}
