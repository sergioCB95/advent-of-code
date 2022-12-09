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

func checkLeft(value int, x int, line string) bool {
	for i := 0; i < x; i++ {
		v, _ := strconv.Atoi(string(rune(line[i])))
		if v >= value {
			return false
		}
	}
	return true
}

func checkRight(value int, x int, line string) bool {
	for i := x + 1; i < len(line); i++ {
		v, _ := strconv.Atoi(string(rune(line[i])))
		if v >= value {
			return false
		}
	}
	return true
}

func checkTop(value int, x int, y int, lines []string) bool {
	for i := 0; i < y; i++ {
		v, _ := strconv.Atoi(string(rune(lines[i][x])))
		if v >= value {
			return false
		}
	}
	return true
}

func checkBottom(value int, x int, y int, lines []string) bool {
	for i := y + 1; i < len(lines); i++ {
		v, _ := strconv.Atoi(string(rune(lines[i][x])))
		if v >= value {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	grid := getFileAsString("../data.txt")
	lines := strings.Split(grid, "\n")
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}

		for x, r := range line {
			value, _ := strconv.Atoi(string(rune(r)))
			if checkLeft(value, x, line) ||
				checkRight(value, x, line) ||
				checkTop(value, x, y, lines) ||
				checkBottom(value, x, y, lines) {
				sum += 1
			}
		}
	}
	println(sum)
}
