package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

		var firstByte, lastByte byte
		for i := 0; i < len(t); i++ {
			if isDigit(t[i]) {
				firstByte = t[i]
				break
			}
		}
		for i := len(t) - 1; i > -1; i-- {
			if isDigit(t[i]) {
				lastByte = t[i]
				break
			}
		}

		num := string(firstByte) + string(lastByte)
		conv, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}

		total += conv
	}
	return total, nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
