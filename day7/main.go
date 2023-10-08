package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Files map[string]int
type Dirs map[string]*Dir
type Dir struct {
	files Files
	dirs  Dirs
	size  int
}

func getDirsUnderLimitTotalSize(d *Dir, limit int) int {
	var size int
	if d.size <= limit {
		size = d.size
	}
	for _, d := range d.dirs {
		s := getDirsUnderLimitTotalSize(d, limit)
		size += s

	}
	return size
}

func findSmallestSubDirToDelete(dir *Dir, minSize int) *Dir {
	var dirToDelete *Dir
	if dirToDelete == nil || (dir.size > minSize && dir.size < dirToDelete.size) {
		dirToDelete = dir
	}

	for _, subDir := range dir.dirs {
		subDirToDelete := findSmallestSubDirToDelete(subDir, minSize)
		if subDirToDelete.size > minSize && subDirToDelete.size < dirToDelete.size {
			dirToDelete = subDirToDelete
		}
	}
	return dirToDelete
}

func newDirectory() *Dir {
	return &Dir{
		files: make(Files),
		dirs:  make(Dirs),
	}
}

const (
	maxDirSize    = 100_000
	totalSpace    = 70_000_000
	requiredSpace = 30_000_000
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	path := []*Dir{newDirectory()}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			cmd := strings.Fields(line)[1:]
			if cmd[0] == "cd" {
				if cmd[1] == ".." {
					path = path[:len(path)-1]
				} else {
					dirName := cmd[1]
					currDir := path[len(path)-1]
					d, ok := currDir.dirs[dirName]
					if !ok {
						dir := newDirectory()
						currDir.dirs[dirName] = dir
						path = append(path, dir)
						continue
					}
					path = append(path, d)
				}
			}
		} else {
			fields := strings.Fields(line)
			currDir := path[len(path)-1]
			if fields[0] == "dir" {
				if _, ok := currDir.dirs[fields[1]]; !ok {
					currDir.dirs[fields[1]] = newDirectory()
				}
			} else {
				size, err := strconv.Atoi(fields[0])
				if err != nil {
					log.Fatalf("Error converting file size: %+v", err)
				}
				currDir.files[fields[1]] = size
				for _, d := range path {
					d.size += size
				}
			}
		}
	}
	f.Close()

	rootDir := path[0]

	dirsUnder100kTotalSize := getDirsUnderLimitTotalSize(rootDir, maxDirSize)
	spaceNeeded := requiredSpace - (totalSpace - rootDir.size)

	dirToDelete := findSmallestSubDirToDelete(rootDir, spaceNeeded)

	fmt.Printf("Part1: %d\n", dirsUnder100kTotalSize)
	fmt.Printf("Part2: %d\n", dirToDelete.size)
}
