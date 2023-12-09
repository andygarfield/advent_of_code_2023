package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type node struct {
	value string
	left  *node
	right *node
}

func main() {
	f, err := os.Open("08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	directions, startNode := parseInput(f)
	steps := solve(directions, startNode)
	fmt.Println(steps)
}

func solve(directions []rune, startNode node) int {
	currentNode := startNode
	directionIndex := 0

	steps := 0
	for currentNode.value != "ZZZ" {
		if directionIndex == len(directions) {
			directionIndex = 0
		}
		direction := directions[directionIndex]

		if direction == 'L' {
			currentNode = *currentNode.left
		} else {
			currentNode = *currentNode.right
		}

		steps++
		directionIndex++
	}

	return steps
}

func parseInput(r io.ReadSeeker) ([]rune, node) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	directions := scanner.Text()

	// blank line
	scanner.Scan()

	// create nodes
	nodes := []node{}
	indexMap := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()
		value := line[:strings.Index(line, " ")]
		nodes = append(nodes, node{value: value})
		indexMap[value] = len(nodes) - 1
	}

	// link nodes
	r.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(r)
	// skip first two lines
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		value := line[:strings.Index(line, " ")]

		startParens := strings.Index(line, "(")
		endParens := strings.Index(line, ")")
		lr := strings.Split(line[startParens+1:endParens], ", ")

		targetIndex := indexMap[value]
		target := &nodes[targetIndex]
		leftNode := &nodes[indexMap[lr[0]]]
		rightNode := &nodes[indexMap[lr[1]]]
		target.left = leftNode
		target.right = rightNode
	}

	return []rune(directions), nodes[indexMap["AAA"]]
}
