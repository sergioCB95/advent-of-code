package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Operation int

const (
	Sum Operation = iota
	Mult
)

type Monkey struct {
	items               [][]int
	operation           Operation
	increment           int
	isVariableIncrement bool
	test                int
	trueMonkeyIndex     int
	falseMonkeyIndex    int
	inspectCount        int
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseItemList(line string, modulusList []int) (itemList [][]int) {
	startingItems := strings.Split(line, ":")
	items := strings.Split(startingItems[1], ",")
	for _, item := range items {
		cleanItem, _ := strconv.Atoi(strings.Trim(item, " "))
		var itemModulus []int
		for _, modulus := range modulusList {
			itemModulus = append(itemModulus, cleanItem%modulus)
		}
		itemList = append(itemList, itemModulus)
	}
	return
}

func mapOperation(s string) Operation {
	switch s {
	case "*":
		return Mult
	case "+":
		return Sum
	}
	return Sum
}

func parseOperation(line string) (operation Operation, increment int, isVariableIncrement bool) {
	operationLine := strings.Split(line, "=")
	operationArgs := strings.Split(strings.Trim(operationLine[1], " "), " ")
	operation = mapOperation(operationArgs[1])
	if operationArgs[2] == "old" {
		isVariableIncrement = true
		increment = 0
	} else {
		isVariableIncrement = false
		increment, _ = strconv.Atoi(operationArgs[2])
	}

	return
}

func parseTest(line string) (test int) {
	testLine := strings.Split(line, " ")
	test, _ = strconv.Atoi(testLine[len(testLine)-1])
	return
}

func parseMonkey(monkeyInput string, modulusList []int) (newMonkey Monkey) {
	lines := strings.Split(monkeyInput, "\n")
	newMonkey.items = parseItemList(lines[1], modulusList)
	newMonkey.operation, newMonkey.increment, newMonkey.isVariableIncrement = parseOperation(lines[2])
	newMonkey.test = parseTest(lines[3])
	newMonkey.trueMonkeyIndex = parseTest(lines[4])
	newMonkey.falseMonkeyIndex = parseTest(lines[5])
	return
}

func main() {
	var monkeys []Monkey
	var modulusList []int
	input := getFileAsString("../data.txt")
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "Test: divisible by") {
			words := strings.Split(line, " ")
			modulus, _ := strconv.Atoi(words[len(words)-1])
			modulusList = append(modulusList, modulus)
		}
	}
	if modulusList != nil {
		monkeyInputs := strings.Split(input, "\n\n")
		for _, monkeyInput := range monkeyInputs {
			monkeys = append(monkeys, parseMonkey(monkeyInput, modulusList))
		}
		for i := 0; i < 10000; i++ {
			for x, monkey := range monkeys {
				for _, item := range monkey.items {
					worryLevels := item
					for j, worryLevel := range worryLevels {
						var increment int
						if monkey.isVariableIncrement {
							increment = worryLevel
						} else {
							increment = monkey.increment
						}
						if monkey.operation == Sum {
							worryLevels[j] = (worryLevels[j] + increment) % modulusList[j]
						} else if monkey.operation == Mult {
							worryLevels[j] = (worryLevels[j] * increment) % modulusList[j]
						}
					}
					if worryLevels[x] == 0 {
						monkeys[monkey.trueMonkeyIndex].items = append(monkeys[monkey.trueMonkeyIndex].items, worryLevels)
					} else {
						monkeys[monkey.falseMonkeyIndex].items = append(monkeys[monkey.falseMonkeyIndex].items, worryLevels)
					}
					monkeys[x].inspectCount++
				}
				monkeys[x].items = [][]int{}
			}
		}
		fmt.Println(modulusList)
		fmt.Println(monkeys)
		for _, monkey := range monkeys {
			println(monkey.inspectCount)
		}
	}
}
