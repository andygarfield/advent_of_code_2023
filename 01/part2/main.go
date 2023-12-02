package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	f, err := os.Open("01/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total, err := findTotal(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}

func findTotal(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		t := scanner.Text()
		digits := extractDigits(t)

		conv, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		if err != nil {
			return 0, err
		}

		total += conv
	}
	return total, nil
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func extractDigits(line string) string {
	buffer := ""
	digits := ""
	for i := 0; i < len(line); i++ {
		buffer += string(line[i])

		if isDigit(line[i]) {
			digits += string(line[i])
		} else {
			for key := range numMap {
				if strings.HasSuffix(buffer, key) {
					digits += numMap[buffer[len(buffer)-len(key):]]
					break
				}
			}
		}
	}

	return digits
}
