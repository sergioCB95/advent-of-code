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

func moveDirCommand(args []string, dirs []*Dir, insideDirs []*Dir) (newDirs []*Dir, newInsideDirs []*Dir) {
	newDirs = dirs
	newInsideDirs = insideDirs
	if args[1] == "cd" {
		if args[2] == ".." {
			if len(newInsideDirs) > 0 {
				newInsideDirs = append([]*Dir{}, newInsideDirs[:len(newInsideDirs)-1]...)
			}
		} else {
			isNewDir := true
			for _, dir := range newDirs {
				if dir.name == getFullName(args[2], newInsideDirs) {
					isNewDir = false
					newInsideDirs = append(newInsideDirs, dir)
				}
			}
			if isNewDir {
				name := getFullName(args[2], newInsideDirs)
				dir := Dir{name: name, size: 0}
				newDirs = append(newDirs, &dir)
				newInsideDirs = append(newInsideDirs, &dir)
			}
		}
	}
	return newDirs, newInsideDirs
}

func addFileSize(fileSize string, insideDirs []*Dir) {
	fs, _ := strconv.Atoi(fileSize)
	for _, dir := range insideDirs {
		dir.size += fs
	}
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
			dirs, insideDirs = moveDirCommand(args, dirs, insideDirs)
		} else if args[0] != "dir" {
			addFileSize(args[0], insideDirs)
		}
	}
	for _, dir := range dirs {
		if dir.size <= 100000 {
			sum += dir.size
		}
	}
	println(sum)
}
