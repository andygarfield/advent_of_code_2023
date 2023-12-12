package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCharGoes(t *testing.T) {
	testCases := []struct {
		c      rune
		d      direction
		result bool
	}{
		{'-', north, false},
		{'-', south, false},
		{'-', east, true},
		{'-', west, true},

		{'|', north, true},
		{'|', south, true},
		{'|', east, false},
		{'|', west, false},

		{'J', north, true},
		{'J', south, false},
		{'J', east, false},
		{'J', west, true},

		{'L', north, true},
		{'L', south, false},
		{'L', east, true},
		{'L', west, false},

		{'7', north, false},
		{'7', south, true},
		{'7', east, false},
		{'7', west, true},

		{'F', north, false},
		{'F', south, true},
		{'F', east, true},
		{'F', west, false},
	}
	for _, testCase := range testCases {
		result := charGoes(testCase.c, testCase.d)
		if result != testCase.result {
			t.Errorf("%v %v %v", testCase.c, testCase.d, testCase.result)
		}
	}
}

func TestFindAdjacentPipes(t *testing.T) {
	testCases := []struct {
		filename string
		start    pos
		pos0     pos
		pos1     *pos
	}{
		{"1.txt", pos{x: 1, y: 1}, pos{x: 1, y: 0}, &pos{x: 1, y: 2}},
		{"2.txt", pos{x: 0, y: 2}, pos{x: 0, y: 3}, &pos{x: 1, y: 2}},
		{"2.txt", pos{x: 2, y: 3}, pos{x: 3, y: 3}, &pos{x: 1, y: 3}},
		{"2.txt", pos{x: 3, y: 0}, pos{x: 3, y: 1}, &pos{x: 2, y: 0}},
	}

	for _, testCase := range testCases {
		input, err := os.ReadFile(fmt.Sprintf("test_data/%s", testCase.filename))
		if err != nil {
			t.Error(err)
		}
		_, inputArea := parseInput(input)
		adj := findAdjacentPipes(inputArea, testCase.start, nil)

		if adj[0].x != testCase.pos0.x || adj[0].y != testCase.pos0.y {
			t.Errorf("adj[0]: %#v does not match expected: %#v", adj[0], testCase.pos0)
		}
		if testCase.pos1 != nil {
			if adj[1].x != testCase.pos1.x || adj[1].y != testCase.pos1.y {
				t.Errorf("adj[1]: %#v does not match expected: %#v", adj[1], testCase.pos1)
			}
		}
	}
}

func TestSolve(t *testing.T) {
	testCases := []struct {
		filename         string
		expectedMaxDist  int
		expectedEnclosed int
	}{
		{"3.txt", 80, 10},
		{"4.txt", 23, 4},
	}
	for _, testCase := range testCases {
		b, err := os.ReadFile(fmt.Sprintf("test_data/%s", testCase.filename))
		if err != nil {
			t.Error(err)
		}
		start, inputArea := parseInput(b)
		replaceStartCharacter(start, inputArea)
		maxDist, numberInside := solve(start, inputArea)
		if maxDist != testCase.expectedMaxDist {
			t.Fail()
		}
		if numberInside != testCase.expectedEnclosed {
			t.Fail()
		}
	}
}
