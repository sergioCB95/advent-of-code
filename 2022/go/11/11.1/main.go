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
	items               []int
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

func parseItemList(line string) (itemList []int) {
	startingItems := strings.Split(line, ":")
	items := strings.Split(startingItems[1], ",")
	for _, item := range items {
		cleanItem, _ := strconv.Atoi(strings.Trim(item, " "))
		itemList = append(itemList, cleanItem)
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

func parseMonkey(monkeyInput string) (newMonkey Monkey) {
	lines := strings.Split(monkeyInput, "\n")
	newMonkey.items = parseItemList(lines[1])
	newMonkey.operation, newMonkey.increment, newMonkey.isVariableIncrement = parseOperation(lines[2])
	newMonkey.test = parseTest(lines[3])
	newMonkey.trueMonkeyIndex = parseTest(lines[4])
	newMonkey.falseMonkeyIndex = parseTest(lines[5])
	return
}

func main() {
	var monkeys []Monkey
	input := getFileAsString("../data.txt")
	monkeyInputs := strings.Split(input, "\n\n")
	for _, monkeyInput := range monkeyInputs {
		monkeys = append(monkeys, parseMonkey(monkeyInput))
	}
	for i := 0; i < 20; i++ {
		for x, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := item
				var increment int
				if monkey.isVariableIncrement {
					increment = worryLevel
				} else {
					increment = monkey.increment
				}
				if monkey.operation == Sum {
					worryLevel += increment
				} else if monkey.operation == Mult {
					worryLevel *= increment
				}
				worryLevel = worryLevel / 3
				if worryLevel%monkey.test == 0 {
					monkeys[monkey.trueMonkeyIndex].items = append(monkeys[monkey.trueMonkeyIndex].items, worryLevel)
				} else {
					monkeys[monkey.falseMonkeyIndex].items = append(monkeys[monkey.falseMonkeyIndex].items, worryLevel)
				}
				monkeys[x].inspectCount++
			}
			monkeys[x].items = []int{}
		}
		fmt.Println(monkeys)
	}
}
