package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileAsString (fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	max := 0
	input := getFileAsString("../data.txt")
	elves := strings.Split(input, "\n\n")
	for _, elf := range elves {
		sum := 0
		calories := strings.Split(elf, "\n")
		for _, calory := range calories {
			num, _ := strconv.Atoi(calory)
			sum += num
		}
		if sum > max {
			max = sum
		}
	}
	println(max)
}

