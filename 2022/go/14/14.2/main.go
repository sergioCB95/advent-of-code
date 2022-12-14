package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  Coord
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data Coord) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data Coord) {
	if n == nil {
		return
	} else if data.x < n.data.x {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else if data.x == n.data.x {
		if data.y <= n.data.y {
			if n.left == nil {
				n.left = &BinaryNode{data: data, left: nil, right: nil}
			} else {
				n.left.insert(data)
			}
		} else {
			if n.right == nil {
				n.right = &BinaryNode{data: data, left: nil, right: nil}
			} else {
				n.right.insert(data)
			}
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) search(data Coord) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.search(data)
	}
}

func (n *BinaryNode) search(data Coord) bool {
	if n == nil {
		return false
	} else if data.x < n.data.x {
		return n.left.search(data)
	} else if data.x == n.data.x {
		if data.y < n.data.y {
			return n.left.search(data)
		} else if data.y > n.data.y {
			return n.right.search(data)
		} else {
			return true
		}
	} else {
		return n.right.search(data)
	}
}

type Coord struct {
	x int
	y int
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getRockStructureCoords(input string) (rockStructureCoordsList [][]Coord) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var rockStructureCoords []Coord
		rawCoords := strings.Split(line, "->")
		for _, rawCoord := range rawCoords {
			coord := strings.Split(strings.Trim(rawCoord, " "), ",")
			parsedX, _ := strconv.Atoi(coord[0])
			parsedY, _ := strconv.Atoi(coord[1])
			rockStructureCoords = append(rockStructureCoords, Coord{x: parsedX, y: parsedY})
		}
		rockStructureCoordsList = append(rockStructureCoordsList, rockStructureCoords)
	}
	return
}

func fillRocks(rockStructureCoordsList [][]Coord) *BinaryTree {
	rockList := &BinaryTree{}
	for _, rockStructureCoords := range rockStructureCoordsList {
		for i, _ := range rockStructureCoords {
			if i == 0 {
				continue
			}
			coord1 := rockStructureCoords[i-1]
			coord2 := rockStructureCoords[i]
			if coord1.x == coord2.x {
				var small, big int
				if coord1.y < coord2.y {
					small = coord1.y
					big = coord2.y
				} else {
					small = coord2.y
					big = coord1.y
				}
				for j := small; j <= big; j++ {
					rockList.insert(Coord{x: coord1.x, y: j})
				}
			} else if coord1.y == coord2.y {
				var small, big int
				if coord1.x < coord2.x {
					small = coord1.x
					big = coord2.x
				} else {
					small = coord2.x
					big = coord1.x
				}
				for j := small; j <= big; j++ {
					rockList.insert(Coord{x: j, y: coord1.y})
				}
			}
		}
	}
	return rockList
}

func main() {
	sum := 0
	var maxY int
	input := getFileAsString("../data.txt")
	rockStructureCoordsList := getRockStructureCoords(input)
	for _, rockStructureCoords := range rockStructureCoordsList {
		for _, coord := range rockStructureCoords {
			if coord.y > maxY {
				maxY = coord.y
			}

		}
	}
	maxY += 2
	notEmptyList := fillRocks(rockStructureCoordsList)
	fall := false
	for !fall {
		pos := Coord{x: 500, y: 0}
		stuck := false
		for !(stuck || fall) {
			if maxY == pos.y+1 {
				sum++
				stuck = true
				notEmptyList.insert(pos)
			}
			if !notEmptyList.search(Coord{x: pos.x, y: pos.y + 1}) {
				pos = Coord{x: pos.x, y: pos.y + 1}
			} else if !notEmptyList.search(Coord{x: pos.x - 1, y: pos.y + 1}) {
				pos = Coord{x: pos.x - 1, y: pos.y + 1}
			} else if !notEmptyList.search(Coord{x: pos.x + 1, y: pos.y + 1}) {
				pos = Coord{x: pos.x + 1, y: pos.y + 1}
			} else {
				sum++
				stuck = true
				notEmptyList.insert(pos)
				if pos.x == 500 && pos.y == 0 {
					fall = true
				}
			}
		}
	}
	println(sum)
}
