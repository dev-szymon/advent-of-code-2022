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

func getTreeHeight(r rune) int {
	if r == 'S' {
		return int('a')
	} else if r == 'E' {
		return int('z')
	} else {
		return int(r)
	}
}

type Node struct {
	y                  int
	x                  int
	value              int
	validAdjacentNodes map[string]*Node
}

func walk(start *Node, end *Node, currentNode *Node, distances map[*Node]float64, visited map[*Node]bool) map[*Node]float64 {
	if currentNode == end {
		return distances
	}

	for _, n := range currentNode.validAdjacentNodes {
		distanceToStart := distances[currentNode] + 1
		if distanceToStart < distances[n] {
			distances[n] = distanceToStart
		}
	}
	visited[currentNode] = true

	var nextNode *Node
	for node, distance := range distances {
		if nextNode == nil {
			nextNode = node
		} else if !visited[node] && distance < distances[nextNode] {
			nextNode = node
		}
	}
	if nextNode == nil {
		return distances
	}

	return walk(start, end, nextNode, distances, visited)
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v\n", err)
	}

	var start *Node
	var end *Node
	matrix := [][]*Node{}

	for y, row := range strings.Split(string(f), "\n") {
		matrix = append(matrix, []*Node{})
		for x, col := range row {
			node := &Node{
				y:                  y,
				x:                  x,
				value:              getTreeHeight(col),
				validAdjacentNodes: make(map[string]*Node),
			}
			if col == 'S' {
				start = node
			} else if col == 'E' {
				end = node
			}

			matrix[y] = append(matrix[y], node)
		}
	}

	visitedNodes := map[*Node]bool{}
	nodesToDistance := map[*Node]float64{}
	for y, row := range matrix {
		for x, node := range row {
			for d, offset := range adjacentPositions {
				adjacentY := y + offset[0]
				adjacentX := x + offset[1]

				if adjacentY >= 0 && adjacentY < len(matrix) && adjacentX >= 0 && adjacentX < len(matrix[y]) {
					adjacentNode := matrix[adjacentY][adjacentX]
					if adjacentNode.value <= node.value+1 {
						node.validAdjacentNodes[d] = adjacentNode
					}
				}
			}
			visitedNodes[node] = false
			nodesToDistance[node] = math.Inf(1)
		}
	}

	nodesToDistance[start] = 0

	distances := walk(start, end, start, nodesToDistance, visitedNodes)

	fmt.Println(distances[end])
}
