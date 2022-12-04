package main

import (
	"io/ioutil"
	"strings"
)

func getFileAsString (fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func mapCharToValue (s rune) int {
	if strings.ToUpper(string(s)) == string(s) {
		return int(s) - int('A') + 27
	}
	return int(s) - int('a') + 1
}

func main() {
	sum := 0
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var commonChars []rune
		firstHalf := line[:len(line) / 2]
		secondHalf := line[len(line) / 2:]
		for _, x := range firstHalf {
			for _, y := range secondHalf {
				if x == y {
					commonChars = append(commonChars, x)
				}
			}
		}
		sum += mapCharToValue(commonChars[0])
	}
	println(sum)
}

