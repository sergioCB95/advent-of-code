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
	sum := 0
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sections := strings.Split(line, ",")
		section1 := strings.Split(sections[0], "-")
		section1Min, _ := strconv.Atoi(section1[0])
		section1Max, _ := strconv.Atoi(section1[1])
		section2 := strings.Split(sections[1], "-")
		section2Min, _ := strconv.Atoi(section2[0])
		section2Max, _ := strconv.Atoi(section2[1])

		if (section1Max >= section2Min && section1Min <= section2Min) || (section2Max >= section1Min && section2Min <= section1Min) {
			sum += 1
		}
	}
	println(sum)
}

