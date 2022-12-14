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

func printGrid(grid [][]Space) {
	for _, row := range grid {
		for _, r := range row {
			print(string(r))
		}
		println()
	}
}

func main() {
	sum := 0
	var maxX, maxY int
	var grid [][]Space
	input := getFileAsString("../data.txt")
	rockStructureCoordsList := getRockStructureCoords(input)

	for _, rockStructureCoords := range rockStructureCoordsList {
		for _, coord := range rockStructureCoords {
			if coord.x > maxX {
				maxX = coord.x
			}
			if coord.y > maxY {
				maxY = coord.y
			}

		}
	}
	for i := 0; i < maxY+1; i++ {
		var subArray []Space
		for j := 0; j < maxX+1; j++ {
			subArray = append(subArray, Empty)
		}
		grid = append(grid, subArray)
	}
	if grid != nil {
		grid = fillGrid(rockStructureCoordsList, grid)
		fall := false
		for !fall {
			pos := Coord{x: 500, y: 0}
			stuck := false
			for !(stuck || fall) {
				if pos.y+1 < len(grid) && grid[pos.y+1][pos.x] == Empty {
					pos = Coord{x: pos.x, y: pos.y + 1}
				} else if pos.y+1 < len(grid) && pos.x-1 >= 0 && grid[pos.y+1][pos.x-1] == Empty {
					pos = Coord{x: pos.x - 1, y: pos.y + 1}
				} else if pos.y+1 < len(grid) && pos.x+1 < len(grid[0]) && grid[pos.y+1][pos.x+1] == Empty {
					pos = Coord{x: pos.x + 1, y: pos.y + 1}
				} else if pos.y+1 >= len(grid) || pos.x-1 < 0 || pos.x+1 >= len(grid[0]) {
					fall = true
				} else {
					sum++
					stuck = true
					grid[pos.y][pos.x] = Sand
				}
			}
		}
	}
	printGrid(grid)
	println(sum)
}
