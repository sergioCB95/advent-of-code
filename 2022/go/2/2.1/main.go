package main

import (
	"io/ioutil"
	"strings"
)

type Result int
const (
	Draw Result = 3
	Lost		= 0
	Win			= 6
)

type Hand interface {
	CheckResult(h Hand) Result
	GetValue() int
}

type Rock struct {}

func (r Rock) CheckResult(h Hand) Result {
	switch h.(type) {
	case Paper:
		return Lost
	case Scissor:
		return Win
	default:
		return Draw
	}
}

func (r Rock) GetValue() int {
	return 1
}

type Paper struct {}

func (r Paper) CheckResult(h Hand) Result {
	switch h.(type) {
	case Scissor:
		return Lost
	case Rock:
		return Win
	default:
		return Draw
	}
}

func (r Paper) GetValue() int {
	return 2
}

type Scissor struct {}

func (r Scissor) CheckResult(h Hand) Result {
	switch h.(type) {
	case Rock:
		return Lost
	case Paper:
		return Win
	default:
		return Draw
	}
}

func (r Scissor) GetValue() int {
	return 3
}

func mapOpponentHand (input string) Hand {
	var result Hand = nil
	switch input {
	case "A":
		result = Rock{}
	case "B":
		result = Paper{}
	case "C":
		result = Scissor{}
	}
	return result
}

func mapMyHand (input string) Hand {
	var result Hand = nil
	switch input {
	case "X":
		result = Rock{}
	case "Y":
		result = Paper{}
	case "Z":
		result = Scissor{}
	}
	return result
}

func getFileAsString (fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	sum := 0
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		values := strings.Split(line, " ")
		opponent := mapOpponentHand(values[0])
		me := mapMyHand(values[1])
		sum += me.GetValue() + int(me.CheckResult(opponent))
	}
	println(sum)
}

