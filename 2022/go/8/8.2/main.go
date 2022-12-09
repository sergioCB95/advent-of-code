package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getScenicLeft(value int, x int, line string) (result int) {
	for i := x - 1; i >= 0; i-- {
		result += 1
		v, _ := strconv.Atoi(string(rune(line[i])))
		if v >= value {
			break
		}
	}
	return
}

func getScenicRight(value int, x int, line string) (result int) {
	for i := x + 1; i < len(line); i++ {
		result += 1
		v, _ := strconv.Atoi(string(rune(line[i])))
		if v >= value {
			break
		}
	}
	return
}

func getScenicTop(value int, x int, y int, lines []string) (result int) {
	for i := y - 1; i >= 0; i-- {
		result += 1
		v, _ := strconv.Atoi(string(rune(lines[i][x])))
		if v >= value {
			break
		}
	}
	return
}

func getScenicBottom(value int, x int, y int, lines []string) (result int) {
	for i := y + 1; i < len(lines); i++ {
		result += 1
		v, _ := strconv.Atoi(string(rune(lines[i][x])))
		if v >= value {
			break
		}
	}
	return
}

func main() {
	max := 0
	grid := getFileAsString("../data.txt")
	lines := strings.Split(grid, "\n")
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}

		for x, r := range line {
			value, _ := strconv.Atoi(string(rune(r)))
			scenic := getScenicLeft(value, x, line) *
				getScenicRight(value, x, line) *
				getScenicTop(value, x, y, lines) *
				getScenicBottom(value, x, y, lines)
			if scenic > max {
				println(getScenicLeft(value, x, line),
					getScenicRight(value, x, line),
					getScenicTop(value, x, y, lines),
					getScenicBottom(value, x, y, lines))
				println(x, y, scenic)
				max = scenic
			}
		}
	}
	println(max)
}
