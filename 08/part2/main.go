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
	startNodes := findStartNodes(nodes)
	nodeSteps := make([]uint64, len(startNodes))

	directionIndex := 0

	for i, currentNode := range startNodes {
		var steps uint64
		for !strings.HasSuffix(currentNode.value, "Z") {
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
		nodeSteps[i] = steps
	}

	return findLCM(nodeSteps)
}

func findLCM(nums []uint64) uint64 {
	multiples := make([]uint64, len(nums))
	copy(multiples, nums)

	sameNums := func() bool {
		var n uint64
		for i, num := range multiples {
			if i == 0 {
				n = num
				continue
			}
			if n != num {
				return false
			}
		}
		return true
	}

	for !sameNums() {
		var smallest uint64
		smallestIndex := 0
		for i, multiple := range multiples {
			if i == 0 {
				smallest = multiple
				continue
			}

			if multiple < smallest {
				smallest = multiple
				smallestIndex = i
			}
		}
		multiples[smallestIndex] += nums[smallestIndex]
	}

	return multiples[0]
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
