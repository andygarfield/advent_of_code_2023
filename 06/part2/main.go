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

	time, record, err := parseFile(f)
	if err != nil {
		log.Fatal(err)
	}

	solution := solve(time, record)
	fmt.Println(solution)
}

func solve(time int, record int) int {
	timesOver := 0
	for timeHeld := 0; timeHeld <= time; timeHeld++ {
		totalTime := (time - timeHeld) * timeHeld
		if totalTime > record {
			timesOver++
		}
	}

	return timesOver
}

func parseFile(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	timeLine := scanner.Text()[11:]

	scanner.Scan()
	recordLine := scanner.Text()[11:]

	time, err := strconv.Atoi(strings.ReplaceAll(timeLine, " ", ""))
	if err != nil {
		return 0, 0, err
	}

	record, err := strconv.Atoi(strings.ReplaceAll(recordLine, " ", ""))
	if err != nil {
		return 0, 0, err
	}

	return time, record, nil
}
