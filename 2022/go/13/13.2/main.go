package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type ValueType int

type Result int

const (
	Single ValueType = iota
	Array
)

const (
	Ordered Result = iota
	NotOrdered
	NotSure
)

type Value struct {
	parent         *Value
	valueType      ValueType
	singleValue    int
	multipleValues []*Value
	flag           bool
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func printValues(pairs *Value) {
	for _, pair := range pairs.multipleValues {
		for _, value := range pair.multipleValues {
			printValue(value)
			println()
		}
		println()
	}
}

func printValue(value *Value) {
	if value.valueType == Array {
		fmt.Print("[ ")
		for _, child := range value.multipleValues {
			printValue(child)
		}
		fmt.Print(" ]")
	} else {
		fmt.Printf(" %v, ", value.singleValue)
	}
}

func parseInput(input string) *Value {
	curr := &Value{parent: nil, valueType: Array}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numFound := false
		num := ""
		for _, r := range line {
			if r == '[' {
				newValue := Value{parent: curr, valueType: Array}
				curr.multipleValues = append(curr.multipleValues, &newValue)
				curr = &newValue
			} else if r == ']' {
				if numFound {
					value, _ := strconv.Atoi(num)
					newValue := Value{parent: curr, valueType: Single, singleValue: value}
					curr.multipleValues = append(curr.multipleValues, &newValue)
					num = ""
					numFound = false
				}
				curr = curr.parent
			} else if r == ',' {
				if numFound {
					value, _ := strconv.Atoi(num)
					newValue := Value{parent: curr, valueType: Single, singleValue: value}
					curr.multipleValues = append(curr.multipleValues, &newValue)
					num = ""
					numFound = false
				}
			} else if _, err := strconv.Atoi(string(r)); err == nil {
				numFound = true
				num += string(r)
			}
		}
	}
	return curr
}

func checkSingleValues(value1 *Value, value2 *Value) (res Result) {
	if value1.singleValue > value2.singleValue {
		println(value1.singleValue, ">", value2.singleValue)
		res = NotOrdered
	} else if value1.singleValue < value2.singleValue {
		println(value1.singleValue, "<", value2.singleValue)
		res = Ordered
	} else {
		res = NotSure
	}
	return
}

func turnIntoArray(value *Value) Value {
	newValue := Value{
		parent:    value.parent,
		valueType: Array,
	}
	auxValue := Value{
		parent:      &newValue,
		valueType:   Single,
		singleValue: value.singleValue,
	}
	newValue.multipleValues = append(newValue.multipleValues, &auxValue)
	return newValue
}

func checkArray(value1 *Value, value2 *Value) Result {
	res := NotSure
	for x, val1 := range value1.multipleValues {
		if x >= len(value2.multipleValues) {
			fmt.Print(value2.multipleValues)
			println("Right side array shorter")
			return NotOrdered
		}
		val2 := value2.multipleValues[x]
		if val1.valueType == Single && val2.valueType == Single {
			res = checkSingleValues(val1, val2)
		} else if val1.valueType == Array && val2.valueType == Array {
			res = checkArray(val1, val2)
		} else {
			if val1.valueType == Single {
				newValue := turnIntoArray(val1)
				res = checkArray(&newValue, val2)
			} else {
				newValue := turnIntoArray(val2)
				res = checkArray(val1, &newValue)
			}
		}
		if res != NotSure {
			return res
		}
	}
	if len(value1.multipleValues) == len(value2.multipleValues) {
		res = NotSure
	} else {
		res = Ordered
	}
	return res
}

func main() {
	sum := 1
	input := getFileAsString("../data.txt")
	root := parseInput(input)
	newValue1 := Value{parent: root, valueType: Single, singleValue: 2}
	newValue2 := Value{parent: root, valueType: Single, singleValue: 6}
	root.multipleValues = append(root.multipleValues, &Value{parent: root, valueType: Array, multipleValues: []*Value{&newValue1}, flag: true})
	root.multipleValues = append(root.multipleValues, &Value{parent: root, valueType: Array, multipleValues: []*Value{&newValue2}, flag: true})
	printValues(root)
	sortedValues := root.multipleValues
	sort.Slice(sortedValues, func(i int, j int) bool {
		result := checkArray(sortedValues[i], sortedValues[j])
		return result == Ordered
	})
	for i, index := range sortedValues {
		if index.flag {
			sum *= i + 1
		}
		printValue(index)
	}
	println(sum)
}
