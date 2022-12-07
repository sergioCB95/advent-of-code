package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Dir struct {
	name string
	size int
}

func getFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getFullName(dirName string, insideDirs []*Dir) string {
	name := ""
	if len(insideDirs) == 0 {
		return dirName
	}
	currDir := insideDirs[len(insideDirs)-1]
	if currDir.name == "/" {
		name = currDir.name + dirName
	} else {
		name = currDir.name + "/" + dirName
	}
	return name
}

func main() {
	sum := 0
	var dirs []*Dir
	var insideDirs []*Dir
	input := getFileAsString("../data.txt")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." {
					if len(insideDirs) > 0 {
						insideDirs = append([]*Dir{}, insideDirs[:len(insideDirs)-1]...)
					}
				} else {
					isNewDir := true
					for _, dir := range dirs {
						if dir.name == getFullName(args[2], insideDirs) {
							isNewDir = false
							insideDirs = append(insideDirs, dir)
						}
					}
					if isNewDir {
						name := getFullName(args[2], insideDirs)
						dir := Dir{name: name, size: 0}
						dirs = append(dirs, &dir)
						insideDirs = append(insideDirs, &dir)
					}
				}
			}
		} else if args[0] != "dir" {
			fileSize, _ := strconv.Atoi(args[0])
			for _, dir := range insideDirs {
				dir.size += fileSize
			}
		}
	}
	for _, dir := range dirs {
		if dir.size <= 100000 {
			sum += dir.size
		}
	}
	println(sum)
}
