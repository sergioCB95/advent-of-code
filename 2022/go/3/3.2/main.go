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
	groups := [][]string{{}}
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if len(groups[len(groups) - 1]) < 3 {
			newSlice := append(groups[len(groups) - 1], line)
			groups[len(groups) - 1] = newSlice
		} else {
			newSlice := []string{line}
			newGroups := append(groups, newSlice)
			groups = newGroups
		}
	}
	for _, group := range groups {
		var commonChars []rune
		for _, x := range group[0] {
			for _, y := range group[1] {
				if x == y {
					for _, z := range group[2] {
						if x == z {
							commonChars = append(commonChars, x)
						}
					}
				}
			}
		}
		sum += mapCharToValue(commonChars[0])
	}
	println(sum)
}

