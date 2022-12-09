package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func mapDirection(d string) Direction {
	switch d {
	case "U":
		return Up
	case "D":
		return Down
	case "R":
		return Right
	case "L":
		return Left
	}
	return Up
}

func moveH(hPos [2]int, dir Direction) (newPos [2]int) {
	newPos = hPos
	switch dir {
	case Up:
		newPos[1] += 1
	case Down:
		newPos[1] -= 1
	case Right:
		newPos[0] += 1
	case Left:
		newPos[0] -= 1
	}
	return
}

func moveT(tPos [2]int, hPos [2]int) (newPos [2]int) {
	newPos = tPos
	xDiff := hPos[0] - tPos[0]
	yDiff := hPos[1] - tPos[1]
	if xDiff == 0 {
		if yDiff > 1 {
			newPos[1] += 1
		} else if yDiff < -1 {
			newPos[1] -= 1
		}
	} else if yDiff == 0 {
		if xDiff > 1 {
			newPos[0] += 1
		} else if xDiff < -1 {
			newPos[0] -= 1
		}
	} else {
		if yDiff > 1 {
			newPos[1] += 1
			if xDiff > 0 {
				newPos[0] += 1
			} else {
				newPos[0] -= 1
			}
		} else if yDiff < -1 {
			newPos[1] -= 1
			if xDiff > 0 {
				newPos[0] += 1
			} else {
				newPos[0] -= 1
			}
		} else if xDiff > 1 {
			newPos[0] += 1
			if yDiff > 0 {
				newPos[1] += 1
			} else {
				newPos[1] -= 1
			}
		} else if xDiff < -1 {
			newPos[0] -= 1
			if yDiff > 0 {
				newPos[1] += 1
			} else {
				newPos[1] -= 1
			}
		}
	}
	return
}

func main() {
	hPos := [2]int{0, 0}
	tPos := [2]int{0, 0}
	var listOfPos [][2]int
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		args := strings.Split(line, " ")
		dir := mapDirection(args[0])
		times, _ := strconv.Atoi(args[1])
		for i := 0; i < times; i++ {
			hPos = moveH(hPos, dir)
			tPos = moveT(tPos, hPos)
			found := false
			for _, pos := range listOfPos {
				if pos[0] == tPos[0] && pos[1] == tPos[1] {
					found = true
				}
			}
			if !found {
				listOfPos = append(listOfPos, tPos)
			}
		}
	}
	println(len(listOfPos))
}
