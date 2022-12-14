package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Space rune

const (
	Empty Space = '.'
	Sand        = 'O'
	Rock        = 'X'
)

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getRockStructureCoords(input string) (rockStructureCoordsList [][]Coord) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var rockStructureCoords []Coord
		rawCoords := strings.Split(line, "->")
		for _, rawCoord := range rawCoords {
			coord := strings.Split(strings.Trim(rawCoord, " "), ",")
			parsedX, _ := strconv.Atoi(coord[0])
			parsedY, _ := strconv.Atoi(coord[1])
			rockStructureCoords = append(rockStructureCoords, Coord{x: parsedX, y: parsedY})
		}
		rockStructureCoordsList = append(rockStructureCoordsList, rockStructureCoords)
	}
	return
}

func fillGrid(rockStructureCoordsList [][]Coord, grid [][]Space) [][]Space {
	for _, rockStructureCoords := range rockStructureCoordsList {
		for i, _ := range rockStructureCoords {
			if i == 0 {
				continue
			}
			coord1 := rockStructureCoords[i-1]
			coord2 := rockStructureCoords[i]
			if coord1.x == coord2.x {
				var small, big int
				if coord1.y < coord2.y {
					small = coord1.y
					big = coord2.y
				} else {
					small = coord2.y
					big = coord1.y
				}
				for j := small; j <= big; j++ {
					grid[j][coord1.x] = Rock
				}
			} else if coord1.y == coord2.y {
				var small, big int
				if coord1.x < coord2.x {
					small = coord1.x
					big = coord2.x
				} else {
					small = coord2.x
					big = coord1.x
				}
				for j := small; j <= big; j++ {
					grid[coord1.y][j] = Rock
				}
			}
		}
	}
	return grid
}

func fillRocks(rockStructureCoordsList [][]Coord) (rockList []Coord) {
	for _, rockStructureCoords := range rockStructureCoordsList {
		for i, _ := range rockStructureCoords {
			if i == 0 {
				continue
			}
			coord1 := rockStructureCoords[i-1]
			coord2 := rockStructureCoords[i]
			if coord1.x == coord2.x {
				var small, big int
				if coord1.y < coord2.y {
					small = coord1.y
					big = coord2.y
				} else {
					small = coord2.y
					big = coord1.y
				}
				for j := small; j <= big; j++ {
					rockList = append(rockList, Coord{x: coord1.x, y: j})
				}
			} else if coord1.y == coord2.y {
				var small, big int
				if coord1.x < coord2.x {
					small = coord1.x
					big = coord2.x
				} else {
					small = coord2.x
					big = coord1.x
				}
				for j := small; j <= big; j++ {
					rockList = append(rockList, Coord{x: j, y: coord1.y})
				}
			}
		}
	}
	return rockList
}

func printGrid(grid [][]Space) {
	for _, row := range grid {
		for _, r := range row {
			print(string(r))
		}
		println()
	}
}

func findCoordInList(x int, y int, coordList []Coord) bool {
	coord := Coord{x: x, y: y}
	for _, c := range coordList {
		if c.x == coord.x && c.y == coord.y {
			return true
		}
	}
	return false
}

func main() {
	sum := 0
	var maxY int
	var grid [][]Space
	var rockList, sandList []Coord
	input := getFileAsString("../data.txt")
	rockStructureCoordsList := getRockStructureCoords(input)
	for _, rockStructureCoords := range rockStructureCoordsList {
		for _, coord := range rockStructureCoords {
			if coord.y > maxY {
				maxY = coord.y
			}

		}
	}
	maxY += 2
	rockList = fillRocks(rockStructureCoordsList)
	fall := false
	for !fall {
		pos := Coord{x: 500, y: 0}
		stuck := false
		for !(stuck || fall) {
			if maxY == pos.y+1 {
				sum++
				stuck = true
				sandList = append(sandList, pos)
			}
			if !findCoordInList(pos.x, pos.y+1, rockList) && !findCoordInList(pos.x, pos.y+1, sandList) {
				pos = Coord{x: pos.x, y: pos.y + 1}
			} else if !findCoordInList(pos.x-1, pos.y+1, rockList) && !findCoordInList(pos.x-1, pos.y+1, sandList) {
				pos = Coord{x: pos.x - 1, y: pos.y + 1}
			} else if !findCoordInList(pos.x+1, pos.y+1, rockList) && !findCoordInList(pos.x+1, pos.y+1, sandList) {
				pos = Coord{x: pos.x + 1, y: pos.y + 1}
			} else {
				sum++
				stuck = true
				sandList = append(sandList, pos)
				if pos.x == 500 && pos.y == 0 {
					fall = true
				}
			}
		}
		println(sum)
	}
	printGrid(grid)
	println(sum)
}
