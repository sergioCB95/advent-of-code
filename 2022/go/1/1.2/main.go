package main

import (
	"io/ioutil"
	"sort"
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
	max3 := []int{ 0, 0, 0 }
	input := getFileAsString("../data.txt")
	elves := strings.Split(input, "\n\n")
	for _, elf := range elves {
		sum := 0
		calories := strings.Split(elf, "\n")
		for _, calory := range calories {
			num, _ := strconv.Atoi(calory)
			sum += num
		}
		if sum > max3[0] {
			max3[0] = sum
			sort.Ints(max3)
		}
	}
	println(max3[0] + max3[1] + max3[2])
}

