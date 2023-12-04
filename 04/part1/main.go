package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		numbersString := line[strings.Index(line, ":")+1:]

		numberSplit := strings.Split(numbersString, "|")
		winningNumberList := numberSplit[0]

		winningNumbers := map[int]struct{}{}
		for _, numString := range strings.Split(winningNumberList, " ") {
			if len(numString) == 0 {
				continue
			}

			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}

			winningNumbers[num] = struct{}{}
		}

		thisCardTotal := 0
		multipler := 0
		for _, numString := range strings.Split(numberSplit[1], " ") {
			if len(numString) == 0 {
				continue
			}

			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}

			if _, exists := winningNumbers[num]; exists {
				thisCardTotal = 1 << multipler
				multipler += 1
			}
		}
		total += thisCardTotal
	}

	fmt.Println(total)
}
