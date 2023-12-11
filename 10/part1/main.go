package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type pos struct {
	x, y int
}

const wait = 50

type direction int

const (
	north direction = iota
	south
	east
	west
)

type area [][]rune

type traverser struct {
	current, previous pos
	area              area
	count             int
}

func main() {
	b, err := os.ReadFile("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start, inputArea := parseInput(b)
	myApp := app.New()
	w := myApp.NewWindow("pipes")
	s := screen{
		pixels:     [980][980]color.RGBA{},
		area:       inputArea,
		traversers: []*traverser{},
		window:     w,
	}
	raster := canvas.NewRaster(func(w, h int) image.Image {
		return &s
	})
	s.raster = raster

	w.Resize(fyne.NewSize(980, 980))
	go func() {
		answer := solve(start, inputArea, &s)
		fmt.Println(answer)
	}()

	w.ShowAndRun()
}

func solve(start pos, inputArea area, thisScreen *screen) int {
	// find loop connected to start
	pipeSides := findAdjacentPipes(inputArea, start, nil)
	t1 := &traverser{current: pipeSides[0], previous: start, area: inputArea, count: 1}
	t2 := &traverser{current: pipeSides[1], previous: start, area: inputArea, count: 1}
	thisScreen.traversers = append(thisScreen.traversers, t1, t2)
	thisScreen.refresh()
	time.Sleep(time.Second)

	t := t1
	for !t1.equals(t2) {
		time.Sleep(wait / 2 * time.Millisecond)

		t.step()
		if t == t1 {
			t = t2
		} else {
			t = t1
		}

		thisScreen.refresh()
		time.Sleep(wait / 2 * time.Millisecond)
	}
	return int(math.Min(float64(t1.count), float64(t2.count)))
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
	t.count++
}

func (t1 *traverser) equals(t2 *traverser) bool {
	return t1.current.x == t2.current.x && t1.current.y == t2.current.y
}

func findAdjacentPipes(a area, start pos, skip *direction) []pos {
	var (
		positions = make([]pos, 2)
		posIndex  = 0
		r         rune
		startChar = a[start.y][start.x]
	)

	// from west side
	if start.x != 0 && (skip == nil || *skip != west) {
		r = a[start.y][start.x-1]
		if (r == '-' || r == 'L' || r == 'F') && (startChar == 'S' || startChar == '-' || startChar == 'J' || startChar == '7') {
			positions[posIndex] = pos{x: start.x - 1, y: start.y}
			posIndex++
		}
	}
	// from east side
	if start.x != len(a[0])-1 && (skip == nil || *skip != east) {
		r = a[start.y][start.x+1]
		if (r == '-' || r == 'J' || r == '7') && (startChar == 'S' || startChar == '-' || startChar == 'L' || startChar == 'F') {
			positions[posIndex] = pos{x: start.x + 1, y: start.y}
			posIndex++
		}
	}
	// from north side
	if start.y != 0 && (skip == nil || *skip != north) {
		r = a[start.y-1][start.x]
		if (r == '|' || r == '7' || r == 'F') && (startChar == 'S' || startChar == '|' || startChar == 'L' || startChar == 'J') {
			positions[posIndex] = pos{x: start.x, y: start.y - 1}
			posIndex++
		}
	}
	// from south side
	if start.y != len(a)-1 && (skip == nil || *skip != south) {
		r = a[start.y+1][start.x]
		if (r == '|' || r == 'L' || r == 'J') && (startChar == 'S' || startChar == '|' || startChar == 'F' || startChar == '7') {
			positions[posIndex] = pos{x: start.x, y: start.y + 1}
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
	var inputArea area = [][]rune{}
	start := pos{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "S") {
			start.x = strings.Index(line, "S")
			start.y = len(inputArea)
		}

		inputArea = append(inputArea, []rune(line))
	}

	return start, inputArea
}
