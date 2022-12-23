package main

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	name     string
	flow     int
	tunnels  []*Valve
	minPaths map[string]int
}

type ValvePairs struct {
	maxPath    int
	timeOpened [2]int
	valves     [2]*Valve
	openedPath []*ValvePairs
}

type ValveDist struct {
	valve   *Valve
	dist    int
	visited bool
	adjs    []*ValveDist
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getValves(input string) (valves []*Valve, workingValves []*Valve, valveDists []*ValveDist) {
	var rawTunnelNames [][]string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		r := regexp.MustCompile(`Valve (?P<name>[A-Z]{2}) has flow rate=(?P<flow>\d+); tunnel[s]? lead[s]? to valve[s]? (?P<rawTunnels>.*)`)
		subStrings := r.FindStringSubmatch(line)
		flow, _ := strconv.Atoi(subStrings[2])
		valve := &Valve{name: subStrings[1], flow: flow}
		valveDist := &ValveDist{valve: valve, visited: false, dist: -1}
		valves = append(valves, valve)
		valveDists = append(valveDists, valveDist)
		if flow > 0 {
			workingValves = append(workingValves, valve)
		}
		tunnelNames := strings.Split(subStrings[3], ", ")
		rawTunnelNames = append(rawTunnelNames, tunnelNames)
	}
	if rawTunnelNames != nil && valveDists != nil {
		for i, _ := range valves {
			for _, name := range rawTunnelNames[i] {
				for j, _ := range valves {
					if name == valves[j].name {
						valves[i].tunnels = append(valves[i].tunnels, valves[j])
						valveDists[i].adjs = append(valveDists[i].adjs, valveDists[j])
					}
				}
			}
		}
	}
	return
}

func minDists(valveDist *ValveDist) map[string]int {
	minPaths := make(map[string]int)
	valveDist.dist = 0
	sortedList := []*ValveDist{valveDist}
	for len(sortedList) > 0 {
		curr := sortedList[0]
		sortedList = sortedList[1:]
		curr.visited = true
		minPaths[curr.valve.name] = curr.dist
		for _, adj := range curr.adjs {
			if !adj.visited {
				if adj.dist < 0 || adj.dist > curr.dist+1 {
					adj.dist = curr.dist + 1
					sortedList = append(sortedList, adj)
				}
			}
		}

		sort.Slice(sortedList, func(i, j int) bool {
			return sortedList[i].dist < sortedList[j].dist
		})
	}
	return minPaths
}

func processValve(curr1 *Valve, time1 int, curr2 *Valve, time2 int, workingValves []*Valve, openedValves []*Valve) int {
	var maxValue int
	var availableOptions1, availableOptions2 []*Valve
	for _, valve := range workingValves {
		alreadyOpened := false
		for _, opened := range openedValves {
			if opened.name == valve.name {
				alreadyOpened = true
			}
		}
		if !alreadyOpened {

			if curr1 != nil {
				newTime1 := time1 - (curr1.minPaths[valve.name] + 1)
				if newTime1 > 0 {
					availableOptions1 = append(availableOptions1, valve)
				}
			}

			if curr2 != nil {
				newTime2 := time2 - (curr2.minPaths[valve.name] + 1)
				if newTime2 > 0 {
					availableOptions2 = append(availableOptions2, valve)
				}
			}
		}
	}
	if len(availableOptions1) > 0 {
		for _, valve1 := range availableOptions1 {
			if len(availableOptions2) > 0 && curr1 != nil && curr2 != nil {
				for _, valve2 := range availableOptions2 {
					if valve1 != valve2 {
						newTime1 := time1 - (curr1.minPaths[valve1.name] + 1)
						newTime2 := time2 - (curr2.minPaths[valve2.name] + 1)
						val := valve1.flow*newTime1 + valve2.flow*newTime2 + processValve(valve1, newTime1, valve2, newTime2, workingValves, append(openedValves, valve1, valve2))
						if val > maxValue {
							maxValue = val
						}
					}
				}
			} else if curr1 != nil {
				newTime1 := time1 - (curr1.minPaths[valve1.name] + 1)
				val := valve1.flow*newTime1 + processValve(valve1, newTime1, nil, 0, workingValves, append(openedValves, valve1))
				if val > maxValue {
					maxValue = val
				}
			}
		}
	} else if curr2 != nil {
		for _, valve2 := range availableOptions2 {
			newTime2 := time2 - (curr2.minPaths[valve2.name] + 1)
			val := valve2.flow*newTime2 + processValve(nil, 0, valve2, newTime2, workingValves, append(openedValves, valve2))
			if val > maxValue {
				maxValue = val
			}
		}
	}

	return maxValue
}

func main() {
	time := 26
	input := getFileAsString("../data.txt")
	var valveA *Valve
	valves, workingValves, valveDists := getValves(input)
	for _, valve := range valves {
		if valve.name == "AA" {
			valveA = valve
		}
	}
	if valveA != nil {
		workingValves = append(workingValves, valveA)
		for _, valve := range workingValves {
			for _, valveDist := range valveDists {
				if valveDist.valve.name == valve.name {
					minPaths := minDists(valveDist)
					valve.minPaths = minPaths
				}
			}
			for _, valveDist := range valveDists {
				valveDist.visited = false
				valveDist.dist = -1
			}
		}
		max := processValve(valveA, time, valveA, time, workingValves, []*Valve{valveA})
		println(max)
	}
}
