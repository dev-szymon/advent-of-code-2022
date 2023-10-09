package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getTreeVision(grid [][]int, y, x int) (int, bool) {
	treeHeight := grid[y][x]

	var (
		upVisionScore    int
		rightVisionScore int
		downVisionScore  int
		leftVisionScore  int
	)
	isUpVisible := true
	isRightVisible := true
	isDownVisibe := true
	isLeftVisible := true

	for up := y - 1; up >= 0; up-- {
		upVisionScore++
		if grid[up][x] >= treeHeight {
			isUpVisible = false
			break
		}
	}

	for right := x + 1; right < len(grid[y]); right++ {
		rightVisionScore++
		if grid[y][right] >= treeHeight {
			isRightVisible = false
			break
		}
	}
	for down := y + 1; down < len(grid); down++ {
		downVisionScore++
		if grid[down][x] >= treeHeight {
			isDownVisibe = false
			break
		}
	}
	for left := x - 1; left >= 0; left-- {
		leftVisionScore++
		if grid[y][left] >= treeHeight {
			isLeftVisible = false
			break
		}
	}
	isOuterTree := y == 0 || x == 0 || y == len(grid)-1 || x == len(grid[y])-1
	isVisible := isOuterTree || isUpVisible || isRightVisible || isDownVisibe || isLeftVisible
	score := upVisionScore * rightVisionScore * downVisionScore * leftVisionScore

	return score, isVisible
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v\n", err)
	}

	rows := strings.Split(string(f), "\n")
	grid := make([][]int, len(rows))

	for y, r := range rows {
		grid[y] = make([]int, len(r))
		cols := strings.Split(r, "")
		for x, c := range cols {
			treeHeight, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("Error while converting tree height to int: %+v\n", err)
			}
			grid[y][x] = treeHeight
		}
	}
	var totalVisible int
	var bestTreeScore int
	for y, row := range grid {
		for x := range row {
			score, isVisible := getTreeVision(grid, y, x)
			if score > bestTreeScore {
				bestTreeScore = score
			}
			if isVisible {
				totalVisible++
			}
		}
	}

	fmt.Printf("Part1: %d\n", totalVisible)
	fmt.Printf("Part2: %d\n", bestTreeScore)
}
