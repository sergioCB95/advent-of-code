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
	name       string
	flow       int
	tunnels    []*Valve
	maxPath    int
	timeOpened int
	minPaths   map[string]int
	openedPath []*Valve
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
		valve := &Valve{name: subStrings[1], flow: flow, maxPath: -1}
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
	time := 30
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
		valveA.maxPath = 0
		valveA.timeOpened = time
		for i := time; i > 0; i-- {
			fmt.Printf("Min Countdown %d\n", i)
			var maxValve *Valve
			var laterValve *Valve
			var maxValue int
			for _, valve := range workingValves {
				if valve.maxPath >= 0 {
					//fmt.Printf("Testing valve %s \n", valve.name)
					for key, value := range valve.minPaths {
						if key != valve.name {
							//println(key, valve.timeOpened - (value + 1), i)
							if valve.timeOpened-(value+1) == i {
								var foundValve *Valve
								for _, x := range workingValves {
									if x.name == key {
										foundValve = x
									}
								}
								if foundValve != nil {
									alreadyOpened := false
									for _, x := range valve.openedPath {
										if x.name == foundValve.name {
											alreadyOpened = true
										}
									}
									if !alreadyOpened && (maxValve == nil || valve.maxPath+(foundValve.flow*i) > maxValue) {
										maxValue = valve.maxPath + (foundValve.flow * i)
										maxValve = foundValve
										laterValve = valve
									}
								}
							}
						}
					}
				}
			}
			if maxValve != nil && laterValve != nil {
				maxValve.maxPath = maxValue
				maxValve.timeOpened = i
				maxValve.openedPath = append(laterValve.openedPath, laterValve)
				fmt.Println(maxValve, laterValve.name)
			}
		}
		fmt.Println(workingValves)
		var max *Valve
		for _, valve := range workingValves {
			if max == nil || max.maxPath < valve.maxPath {
				max = valve
			}
		}
		fmt.Println(max)
	}
}
