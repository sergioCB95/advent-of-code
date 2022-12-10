package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Operation interface {
	increaseRegister(x int) int
	getCompletionCycles() int
}

type Addx struct {
	increment int
}

func (a Addx) increaseRegister(x int) int {
	return x + a.increment
}

func (a Addx) getCompletionCycles() int {
	return 2
}

type Noop struct{}

func (a Noop) increaseRegister(x int) int {
	return x
}

func (a Noop) getCompletionCycles() int {
	return 1
}

func mapOperation(args []string) Operation {
	switch args[0] {
	case "noop":
		return Noop{}
	case "addx":
		increment, _ := strconv.Atoi(args[1])
		return Addx{increment: increment}
	}
	return Noop{}
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	cycle := 0
	xRegister := 1
	screen := [6]string{}
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		operation := mapOperation(strings.Split(line, " "))
		opCycles := operation.getCompletionCycles()
		for i := 0; i < opCycles; i++ {
			cycle++
			hPos := cycle%40 - 1
			vPos := (cycle - 1) / 40
			if hPos >= xRegister-1 && hPos <= xRegister+1 {
				screen[vPos] += "#"
			} else {
				screen[vPos] += "."
			}

		}
		xRegister = operation.increaseRegister(xRegister)
	}
	for _, line := range screen {
		println(line)
	}
}
