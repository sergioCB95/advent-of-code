package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getCargoStacks(cargoInput string) [][]rune {
	cargo := [][]rune{{}}
	for _, line := range strings.Split(cargoInput, "\n") {
		if len(line) == 0 {
			continue
		}
		i := 0
		j := 1 + (4 * i)
		for j < len(line) {
			char := rune(line[j])
			if len(cargo) <= i {
				cargo = append(cargo, []rune{})
			}
			if unicode.IsLetter(char) {
				cargoStack := append(cargo[i], char)
				cargo[i] = cargoStack
			}
			i += 1
			j = 1 + (4 * i)
		}
	}
	return cargo
}

func getActions(actionsInput string) [][]int {
	var actions [][]int
	for _, line := range strings.Split(actionsInput, "\n") {
		if len(line) == 0 {
			continue
		}
		words := strings.Split(line, " ")
		amount, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		to, _ := strconv.Atoi(words[5])
		actions = append(actions, []int{amount, from, to})
	}
	return actions
}

func main() {
	input := getFileAsString("../data.txt")
	parts := strings.Split(input, "\n\n")
	cargoInput := parts[0]
	actionsInput := parts[1]
	cargo := getCargoStacks(cargoInput)
	actions := getActions(actionsInput)

	for _, action := range actions {
		amount := action[0]
		from := action[1]
		to := action[2]
		moved := append([]rune{}, cargo[from-1][:amount]...)
		for _, x := range moved {
			cargo[to-1] = append([]rune{x}, cargo[to-1]...)
		}
		cargo[from-1] = cargo[from-1][amount:]
	}

	for _, x := range cargo {
		for _, y := range x {
			print(strconv.QuoteRune(y))
		}
		println()
	}
}
