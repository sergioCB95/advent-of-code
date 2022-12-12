package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type GridPoint struct {
	value         rune
	xPos          int
	yPos          int
	minPathWeight int
	visited       bool
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func checkAndGetPath(xPos int, yPos int, gridPoint GridPoint, grid [][]GridPoint, sortedList []GridPoint) ([][]GridPoint, []GridPoint) {
	curr := grid[xPos][yPos]
	fmt.Print(curr)
	if !curr.visited && curr.value-gridPoint.value <= 1 && (curr.minPathWeight < 0 || gridPoint.minPathWeight+1 < curr.minPathWeight) {
		newPoint := GridPoint{
			xPos:          xPos,
			yPos:          yPos,
			minPathWeight: gridPoint.minPathWeight + 1,
			value:         curr.value,
		}
		grid[xPos][yPos] = newPoint
		found := false
		for x, item := range sortedList {
			if item.xPos == newPoint.xPos && item.yPos == newPoint.yPos {
				sortedList[x] = newPoint
				found = true
			}
		}
		if !found {
			sortedList = append(sortedList, newPoint)
		}
	}
	return grid, sortedList
}

func main() {
	var grid [][]GridPoint
	var eX, eY int
	var sortedList []GridPoint
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for x, line := range lines {
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []GridPoint{})
		for y, r := range line {
			if r == 'S' {
				first := GridPoint{value: 'a', minPathWeight: 0, xPos: x, yPos: y}
				grid[x] = append(grid[x], first)
				sortedList = append(sortedList, first)
			} else if r == 'E' {
				eX = x
				eY = y
				grid[x] = append(grid[x], GridPoint{value: 'z', minPathWeight: -1, xPos: x, yPos: y})
			} else {
				grid[x] = append(grid[x], GridPoint{value: r, minPathWeight: -1, xPos: x, yPos: y})
			}
		}
	}
	if grid != nil {
		for len(sortedList) > 0 {
			curr := sortedList[0]
			sortedList = sortedList[1:]
			curr.visited = true
			grid[curr.xPos][curr.yPos] = curr

			if curr.xPos == eX && curr.yPos == eY {
				break
			}

			if curr.xPos > 0 {
				grid, sortedList = checkAndGetPath(curr.xPos-1, curr.yPos, curr, grid, sortedList)
			}
			if curr.xPos < (len(grid) - 1) {
				grid, sortedList = checkAndGetPath(curr.xPos+1, curr.yPos, curr, grid, sortedList)
			}
			if curr.yPos > 0 {
				grid, sortedList = checkAndGetPath(curr.xPos, curr.yPos-1, curr, grid, sortedList)
			}
			if curr.yPos < (len(grid[curr.xPos]) - 1) {
				grid, sortedList = checkAndGetPath(curr.xPos, curr.yPos+1, curr, grid, sortedList)
			}

			sort.Slice(sortedList, func(i, j int) bool {
				return sortedList[i].minPathWeight < sortedList[j].minPathWeight
			})
			fmt.Println(sortedList)
		}
		fmt.Println(grid[eX][eY])
	}
}
