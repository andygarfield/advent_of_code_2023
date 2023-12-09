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

	directions, nodes := parseInput(f)
	steps := solve(directions, nodes)
	fmt.Println(steps)
}

func solve(directions []rune, nodes []node) uint64 {
	currentNodes := findStartNodes(nodes)

	directionIndex := 0

	var steps uint64
	for !hasAllZs(currentNodes) {
		if directionIndex == len(directions) {
			directionIndex = 0
		}
		direction := directions[directionIndex]

		for i, currentNode := range currentNodes {
			if direction == 'L' {
				currentNodes[i] = *currentNode.left
			} else {
				currentNodes[i] = *currentNode.right
			}
		}

		steps++
		directionIndex++
	}

	return steps
}

func hasAllZs(nodes []node) bool {
	for _, n := range nodes {
		if !strings.HasSuffix(n.value, "Z") {
			return false
		}
	}
	return true
}

func findStartNodes(nodes []node) []node {
	startNodes := []node{}
	for _, n := range nodes {
		if strings.HasSuffix(n.value, "A") {
			startNodes = append(startNodes, n)
		}
	}

	return startNodes
}

func parseInput(r io.ReadSeeker) ([]rune, []node) {
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

	// create nodes
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

	return []rune(directions), nodes
}
