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
	GetOpponentsHand(result Result) Hand
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

func (r Rock) GetOpponentsHand(result Result) Hand {
	switch result {
	case Win:
		return Paper{}
	case Lost:
		return Scissor{}
	default:
		return Rock{}
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

func (r Paper) GetOpponentsHand(result Result) Hand {
	switch result {
	case Win:
		return Scissor{}
	case Lost:
		return Rock{}
	default:
		return Paper{}
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

func (r Scissor) GetOpponentsHand(result Result) Hand {
	switch result {
	case Win:
		return Rock{}
	case Lost:
		return Paper{}
	default:
		return Scissor{}
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

func mapResult (input string) Result {
	var result Result
	switch input {
	case "X":
		result = Lost
	case "Y":
		result = Draw
	case "Z":
		result = Win
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
		result := mapResult(values[1])
		me := opponent.GetOpponentsHand(result)
		sum += me.GetValue() + int(result)
	}
	println(sum)
}

