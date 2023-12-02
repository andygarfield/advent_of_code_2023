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

type cubeSet struct {
	red, green, blue int
}

type game struct {
	gameID int
	cubes  []cubeSet
}

func main() {
	f, err := os.Open("02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := calculateAnswer(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}

func calculateAnswer(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		id, playable, err := getGameInfo(line)
		if err != nil {
			return 0, err
		}
		if playable {
			total += id
		}
	}

	return total, nil
}

// getGameInfo returns the game id and whether the game can be played
func getGameInfo(line string) (int, bool, error) {
	cc := game{cubes: []cubeSet{}}

	var err error
	colonIndex := strings.Index(line, ":")
	cc.gameID, err = strconv.Atoi(line[5:colonIndex])
	if err != nil {
		return cc.gameID, false, err
	}

	cubeInfo := line[colonIndex+1:]
	sets := strings.Split(cubeInfo, ";")

	for _, set := range sets {
		cubes, err := parseSet(set)
		if err != nil {
			return cc.gameID, false, err
		}
		cc.cubes = append(cc.cubes, cubes)
	}

	return cc.gameID, canBePlayed(cc.cubes), nil
}

func canBePlayed(cubeSlice []cubeSet) bool {
	for _, cubeSet := range cubeSlice {
		if cubeSet.red > 12 {
			return false
		}
		if cubeSet.green > 13 {
			return false
		}
		if cubeSet.blue > 14 {
			return false
		}
	}
	return true
}

func parseSet(s string) (cubeSet, error) {
	cubes := cubeSet{}
	for _, stmt := range strings.Split(s, ",") {
		trimmed := strings.TrimSpace(stmt)
		spaceIndex := strings.Index(trimmed, " ")
		countStr := trimmed[:spaceIndex]
		color := trimmed[spaceIndex+1:]
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return cubes, err
		}
		switch color {
		case "red":
			cubes.red = count
		case "green":
			cubes.green = count
		case "blue":
			cubes.blue = count
		default:
			return cubes, fmt.Errorf(`the color "%s" is not red, green, or blue`, color)
		}
	}
	return cubes, nil
}
