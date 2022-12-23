package main

import (
	"fmt"
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

func main() {
	time := 26
	input := getFileAsString("../data-example.txt")
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
		startingValvePair := ValvePairs{maxPath: 0, timeOpened: [2]int{time, time}, valves: [2]*Valve{valveA, valveA}}
		list := []*ValvePairs{&startingValvePair}
		for i := time; i > 0; i-- {
			fmt.Printf("Min Countdown %d\n", i)
			var maxPairValve *ValvePairs
			var laterPairValve *ValvePairs
			var maxPairValue int
			for _, valvePair := range list {
				var maxValve [2]*Valve
				var maxValue [2]int
				//fmt.Printf("Testing valve %s-%s \n", valvePair.valves[0].name, valvePair.valves[1].name)
				for j, valve := range valvePair.valves {
					for key, value := range valve.minPaths {
						if key != valve.name {
							//println(key, valvePair.timeOpened[j] - (value + 1), i)
							if valvePair.timeOpened[j]-(value+1) == i {
								var foundValve *Valve
								for _, x := range workingValves {
									if x.name == key {
										foundValve = x
									}
								}
								if foundValve != nil {
									alreadyOpened := false
									for _, x := range valvePair.openedPath {
										for _, y := range x.valves {
											if y.name == foundValve.name {
												alreadyOpened = true
											}
										}
									}
									if !alreadyOpened && (j == 0 || maxValve[0] != foundValve) && (maxValve[j] == nil || (foundValve.flow*i) > maxValue[j]) {
										maxValue[j] = foundValve.flow * i
										maxValve[j] = foundValve
									}
								}
							}
						}
					}
				}
				if maxValve[0] != nil || maxValve[1] != nil {
					if maxValue[0]+maxValue[1] > maxPairValue {
						laterPairValve = valvePair
						maxPairValue = maxValue[0] + maxValue[1]
						maxPairValve = &ValvePairs{maxPath: laterPairValve.maxPath + maxPairValue}
						if maxValve[0] != nil {
							maxPairValve.valves[0] = maxValve[0]
							maxPairValve.timeOpened[0] = i
						} else {
							maxPairValve.valves[0] = valvePair.valves[0]
							maxPairValve.timeOpened[0] = valvePair.timeOpened[0]
						}
						if maxValve[1] != nil {
							maxPairValve.valves[1] = maxValve[1]
							maxPairValve.timeOpened[1] = i
						} else {
							maxPairValve.valves[1] = valvePair.valves[1]
							maxPairValve.timeOpened[1] = valvePair.timeOpened[1]
						}
						laterPairValve = valvePair
						maxPairValve.openedPath = append(laterPairValve.openedPath, laterPairValve)
					}
				}
			}
			if maxPairValve != nil {
				fmt.Printf("Selected valve pair %s-%s \n", maxPairValve.valves[0].name, maxPairValve.valves[1].name)
				list = append(list, maxPairValve)
			}
		}
		fmt.Println(workingValves)
		var max *ValvePairs
		for _, valve := range list {
			if max == nil || max.maxPath < valve.maxPath {
				max = valve
			}
		}
		for _, valve := range max.openedPath {
			fmt.Println(valve)
		}
		fmt.Println(max)
	}
}
