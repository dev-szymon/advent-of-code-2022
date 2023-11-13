package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var adjacentPositions = map[string][]int{
	"top":    {-1, 0},
	"right":  {0, 1},
	"bottom": {1, 0},
	"left":   {0, -1},
}

func getElevation(r rune) int {
	if r == 'S' {
		return int('a')
	} else if r == 'E' {
		return int('z')
	} else {
		return int(r)
	}
}

type Point struct {
	y                   int
	x                   int
	elevation           int
	validAdjacentPoints map[string]*Point
}

func walk(start *Point, end *Point, currentPoint *Point, distances map[*Point]float64, visited map[*Point]bool, shortestKnownDistance int) (map[*Point]float64, bool) {
	if shortestKnownDistance > 0 && distances[currentPoint] > float64(shortestKnownDistance) {
		return distances, true
	}

	if currentPoint == end {
		return distances, false
	}

	for _, n := range currentPoint.validAdjacentPoints {
		distanceToStart := distances[currentPoint] + 1
		if distanceToStart < distances[n] {
			distances[n] = distanceToStart
		}
	}
	visited[currentPoint] = true

	var nextNode *Point
	for node, distance := range distances {
		if nextNode == nil {
			nextNode = node
		} else if !visited[node] && distance < distances[nextNode] {
			nextNode = node
		}
	}
	if nextNode == nil {
		return distances, false
	}

	return walk(start, end, nextNode, distances, visited, shortestKnownDistance)
}

func findShortestDistanceBetweenPoints(nodes []*Point, start *Point, end *Point, shortestKnownDistance int) int {
	visitedNodes := map[*Point]bool{}
	nodesToDistance := map[*Point]float64{}

	for _, node := range nodes {
		visitedNodes[node] = false
		if node == start {
			nodesToDistance[node] = 0
		} else {
			nodesToDistance[node] = math.Inf(1)
		}
	}

	distances, terminated := walk(start, end, start, nodesToDistance, visitedNodes, shortestKnownDistance)
	if terminated {
		return -1
	}
	return int(distances[end])
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v\n", err)
	}

	var start *Point
	var end *Point
	part2startingPoints := []*Point{}
	matrix := [][]*Point{}

	for y, row := range strings.Split(string(f), "\n") {
		matrix = append(matrix, []*Point{})
		for x, col := range row {
			e := getElevation(col)
			node := &Point{
				y:                   y,
				x:                   x,
				elevation:           e,
				validAdjacentPoints: make(map[string]*Point),
			}

			if col == 'S' {
				start = node
			} else if col == 'E' {
				end = node
			}

			if e == int('a') {
				part2startingPoints = append(part2startingPoints, node)
			}

			matrix[y] = append(matrix[y], node)
		}
	}

	points := []*Point{}
	for y, row := range matrix {
		for x, point := range row {
			for d, offset := range adjacentPositions {
				adjacentY := y + offset[0]
				adjacentX := x + offset[1]

				if adjacentY >= 0 && adjacentY < len(matrix) && adjacentX >= 0 && adjacentX < len(matrix[y]) {
					adjacentNode := matrix[adjacentY][adjacentX]
					if adjacentNode.elevation <= point.elevation+1 {
						point.validAdjacentPoints[d] = adjacentNode
					}
				}
			}

			points = append(points, point)
		}
	}

	part1 := findShortestDistanceBetweenPoints(points, start, end, -1)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := -1
	for _, pt2start := range part2startingPoints {
		pt2distance := findShortestDistanceBetweenPoints(points, pt2start, end, part1)
		if pt2distance > 0 && (part2 < 0 || pt2distance < part2) {
			part2 = pt2distance
		}
	}
	fmt.Printf("Part 2: %d\n", part2)
}
