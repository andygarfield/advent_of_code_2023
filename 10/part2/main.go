package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

type direction int

const (
	north direction = iota
	south
	east
	west
)

type pipe struct {
	rune    rune
	visited bool
}

type area [][]pipe

type traverser struct {
	current, previous pos
	area              area
	count             int
}

var directionToCharacters = map[direction][]rune{
	north: {'|', 'J', 'L'},
	south: {'|', '7', 'F'},
	east:  {'-', 'L', 'F'},
	west:  {'-', 'J', '7'},
}

var directionOpposites = map[direction]direction{
	north: south,
	south: north,
	east:  west,
	west:  east,
}

func main() {
	b, err := os.ReadFile("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start, inputArea := parseInput(b)
	replaceStartCharacter(start, inputArea)
	_, numberInside := solve(start, inputArea)
	fmt.Println(numberInside)
}

func solve(start pos, inputArea area) (int, int) {
	pipeSides := findAdjacentPipes(inputArea, start, nil)

	inputArea[start.y][start.x].visited = true

	t1 := &traverser{current: pipeSides[0], previous: start, area: inputArea, count: 1}
	t2 := &traverser{current: pipeSides[1], previous: start, area: inputArea, count: 1}
	inputArea[t1.current.y][t1.current.x].visited = true
	inputArea[t2.current.y][t2.current.x].visited = true

	t := t1
	for !t1.equals(t2) {
		t.step()
		if t == t1 {
			t = t2
		} else {
			t = t1
		}
	}

	insideCount := 0
	for y := 0; y < len(inputArea); y++ {
		for x := 0; x < len(inputArea[0]); x++ {
			if inputArea[y][x].visited {
				continue
			}
			leftwardVisitedPipes := 0
			for z := 0; z < x; z++ {
				p := inputArea[y][z]
				if p.visited && (p.rune == 'F' || p.rune == '7' || p.rune == '|') {
					leftwardVisitedPipes++
				}
			}
			if leftwardVisitedPipes%2 == 1 {
				inputArea[y][x].rune = 'I'
				insideCount++
			}
		}
	}

	return int(math.Min(float64(t1.count), float64(t2.count))), insideCount
}

func (t *traverser) step() {
	var from direction
	if t.current.x-1 == t.previous.x {
		from = west
	} else if t.current.x+1 == t.previous.x {
		from = east
	} else if t.current.y-1 == t.previous.y {
		from = north
	} else {
		from = south
	}

	positions := findAdjacentPipes(t.area, t.current, &from)

	t.previous = t.current
	t.current = positions[0]

	t.area[t.current.y][t.current.x].visited = true
	t.count++
}

func (t1 *traverser) equals(t2 *traverser) bool {
	return t1.current.x == t2.current.x && t1.current.y == t2.current.y
}

// findAdjacentPipes finds the pipes connected to the start position in the
// area. Given properly-formed pipe input, this should return a maximum of 2
// items. If one of the directions is already known, providing "skip" as
// non-nil for the side of the start position which is known will make the
// function only return a single-item slice of the remaining 3 sides. The
// ordering will always be north, south, east, then west.
func findAdjacentPipes(a area, start pos, skip *direction) []pos {
	var (
		positions = make([]pos, 2)
		posIndex  = 0
		r         rune
		startChar = a[start.y][start.x].rune
	)
	// from north side
	if start.y != 0 && (skip == nil || *skip != north) {
		r = a[start.y-1][start.x].rune
		if charGoes(r, south) && (charGoes(startChar, north) || startChar == 'S') {
			positions[posIndex] = pos{x: start.x, y: start.y - 1}
			posIndex++
		}
	}
	// from south side
	if start.y != len(a)-1 && (skip == nil || *skip != south) {
		r = a[start.y+1][start.x].rune
		if charGoes(r, north) && (charGoes(startChar, south) || startChar == 'S') {
			positions[posIndex] = pos{x: start.x, y: start.y + 1}
			posIndex++
		}
	}
	// from east side
	if start.x != len(a[0])-1 && (skip == nil || *skip != east) {
		r = a[start.y][start.x+1].rune
		if charGoes(r, west) && (charGoes(startChar, east) || startChar == 'S') {
			positions[posIndex] = pos{x: start.x + 1, y: start.y}
			posIndex++
		}
	}
	// from west side
	if start.x != 0 && (skip == nil || *skip != west) {
		r = a[start.y][start.x-1].rune
		if charGoes(r, east) && (charGoes(startChar, west) || startChar == 'S') {
			positions[posIndex] = pos{x: start.x - 1, y: start.y}
			// no cases left, no need to increment
		}
	}

	if skip == nil {
		return positions
	} else {
		return positions[:1]
	}
}

func parseInput(b []byte) (pos, area) {
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	var inputArea area = [][]pipe{}
	start := pos{}

	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "S") {
			start.x = strings.Index(line, "S")
			start.y = len(inputArea)
		}

		linePipes := []pipe{}
		for _, r := range line {
			linePipes = append(linePipes, pipe{rune: r})
		}
		inputArea = append(inputArea, linePipes)
		lineIndex++
	}

	return start, inputArea
}

// replaceStartCharacter replace S rune in area with the correct pipe character
func replaceStartCharacter(start pos, inputArea area) {
	adjacentToStart := findAdjacentPipes(inputArea, start, nil)

	directions := []direction{}
	for _, adjacentPos := range adjacentToStart {
		if adjacentPos.x < start.x {
			directions = append(directions, west)
		} else if adjacentPos.x > start.x {
			directions = append(directions, east)
		} else if adjacentPos.y < start.y {
			directions = append(directions, north)
		} else if adjacentPos.y > start.y {
			directions = append(directions, south)
		}
	}
	common := ' '
	for _, a := range directionToCharacters[directions[0]] {
		for _, b := range directionToCharacters[directions[1]] {
			if a == b {
				common = a
				break
			}
		}
		if common != ' ' {
			break
		}
	}

	inputArea[start.y][start.x].rune = common
}

func charGoes(c rune, dir direction) bool {
	characters := directionToCharacters[dir]
	for _, mc := range characters {
		if c == mc {
			return true
		}
	}
	return false
}
