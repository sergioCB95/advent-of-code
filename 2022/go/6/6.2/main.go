package main

import (
	"io/ioutil"
	"strings"
)

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	var result int
	input := getFileAsString("../data.txt")
	initI := 0
	endI := 14
	for endI < len(input) {
		found := true
		marker := input[initI:endI]
		println(marker)
		for _, r := range marker {
			count := strings.Count(marker, string(r))
			println(count)
			if count > 1 {
				found = false
			}
		}
		if found {
			result = endI
			break
		}
		println()
		initI += 1
		endI += 1
	}
	println(result)
}
