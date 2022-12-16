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

func getSensorBeaconList(input string) (sensorBeaconList [][]int) {
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
	return
}

func main() {
	maxNum := 4000000
	sum := 0
	borders := make(map[string]int)
	var intersections [][]int
	var solutions [][]int
	input := getFileAsString("../data.txt")
	sensorBeaconList := getSensorBeaconList(input)
	for x, sensorBeacon := range sensorBeaconList {
		distance := math.Abs(float64(sensorBeacon[2])-float64(sensorBeacon[0])) + math.Abs(float64(sensorBeacon[3])-float64(sensorBeacon[1]))
		sensorBeaconList[x] = append(sensorBeacon, int(distance))
	}
	for _, sensorBeacon := range sensorBeaconList {
		distance := sensorBeacon[4] + 1
		for i := 1; i < distance; i++ {
			// down right
			x1 := sensorBeacon[0] + i
			y1 := sensorBeacon[1] + distance - i
			if x1 >= 0 && x1 <= maxNum && y1 >= 0 && y1 <= maxNum {
				key := fmt.Sprintf("%d%v%d", x1, ",", y1)
				borders[key] += 1
			}

			// up right
			x2 := sensorBeacon[0] + i
			y2 := sensorBeacon[1] - (distance - i)
			if x2 >= 0 && x2 <= maxNum && y2 >= 0 && y2 <= maxNum {
				key := fmt.Sprintf("%d%v%d", x2, ",", y2)
				borders[key] += 1
			}

			// down right
			x3 := sensorBeacon[0] - i
			y3 := sensorBeacon[1] + distance - i
			if x3 >= 0 && x3 <= maxNum && y3 >= 0 && y3 <= maxNum {
				key := fmt.Sprintf("%d%v%d", x3, ",", y3)
				borders[key] += 1
			}

			// up left
			x4 := sensorBeacon[0] - i
			y4 := sensorBeacon[1] - (distance - i)
			if x4 >= 0 && x4 <= maxNum && y4 >= 0 && y4 <= maxNum {
				key := fmt.Sprintf("%d%v%d", x3, ",", y3)
				borders[key] += 1
			}
		}
		key1 := fmt.Sprintf("%d%v%d", sensorBeacon[0], ",", sensorBeacon[1]+distance)
		borders[key1] += 1
		key2 := fmt.Sprintf("%d%v%d", sensorBeacon[0], ",", sensorBeacon[1]-distance)
		borders[key2] += 1
		key3 := fmt.Sprintf("%d%v%d", sensorBeacon[0]+distance, ",", sensorBeacon[1])
		borders[key3] += 1
		key4 := fmt.Sprintf("%d%v%d", sensorBeacon[0]-distance, ",", sensorBeacon[1])
		borders[key4] += 1
	}
	for k, v := range borders {
		if v > 3 {
			sum++
			coords := strings.Split(k, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			intersections = append(intersections, []int{x, y})
		}
	}
	for _, intersection := range intersections {
		found := true
		for _, sensorBeacon := range sensorBeaconList {
			distance := math.Abs(float64(intersection[0])-float64(sensorBeacon[0])) + math.Abs(float64(intersection[1])-float64(sensorBeacon[1]))
			if int(distance) <= sensorBeacon[4] {
				found = false
			}
		}
		if found {
			solutions = append(solutions, intersection)
		}
	}
	println(len(borders))
	println(len(intersections))
	fmt.Println(solutions)
}
