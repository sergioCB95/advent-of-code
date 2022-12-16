package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
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
	targetY := 2000000
	resultSet := make(map[int]bool)
	occupiedFields := make(map[int]bool)
	var sensorBeaconList [][]int
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		r := regexp.MustCompile(`Sensor at x=(?P<sx>-?\d+), y=(?P<sy>-?\d+): closest beacon is at x=(?P<bx>-?\d+), y=(?P<by>-?\d+)`)
		subStrings := r.FindStringSubmatch(line)
		var parsedSubstring []int
		for _, s := range subStrings {
			i, _ := strconv.Atoi(s)
			parsedSubstring = append(parsedSubstring, i)
		}
		if parsedSubstring != nil {
			sensorBeaconList = append(sensorBeaconList, parsedSubstring[1:])
		}
	}
	for _, sensorBeacon := range sensorBeaconList {
		if sensorBeacon[1] == targetY {
			occupiedFields[sensorBeacon[0]] = true
		} else if sensorBeacon[3] == targetY {
			occupiedFields[sensorBeacon[2]] = true
		}
	}
	for _, sensorBeacon := range sensorBeaconList {
		distance := math.Abs(float64(sensorBeacon[2])-float64(sensorBeacon[0])) + math.Abs(float64(sensorBeacon[3])-float64(sensorBeacon[1]))
		yDistWithTarget := math.Abs(float64(targetY) - float64(sensorBeacon[1]))
		if diff := int(distance - yDistWithTarget); diff >= 0 {
			for i := 0; i <= diff; i++ {
				if occupiedFields[sensorBeacon[0]-i] != true {
					resultSet[sensorBeacon[0]-i] = true
				}
				if occupiedFields[sensorBeacon[0]+i] != true {
					resultSet[sensorBeacon[0]+i] = true
				}
			}
		}
	}
	fmt.Println(len(resultSet))
}
