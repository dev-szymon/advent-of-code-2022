package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getDirection(direction string) (int, int) {
	switch direction {
	case "U":
		return -1, 0
	case "R":
		return 0, 1
	case "D":
		return 1, 0
	case "L":
		return 0, -1
	default:
		return 0, 0
	}
}

func getDistance(h, t Position) int {
	var diffY int
	var diffX int
	hy := float64(h[0])
	hx := float64(h[1])
	ty := float64(t[0])
	tx := float64(t[1])

	if h[0] < t[0] {
		diffY = int(math.Abs(ty - hy))
	} else {
		diffY = int(math.Abs(hy - ty))
	}
	if h[1] < t[1] {
		diffX = int(math.Abs(tx - hx))
	} else {
		diffX = int(math.Abs(hx - tx))
	}

	if diffX > diffY {
		return diffX
	} else {
		return diffY
	}
}

func move(h, t Position, direction string) (Position, Position) {
	y, x := getDirection(direction)
	var nextHead Position
	nextHead[0] = h[0] + y
	nextHead[1] = h[1] + x

	isAdjacent := getDistance(nextHead, t) < 2

	if isAdjacent {
		return nextHead, t
	} else {
		nextTail := h
		return nextHead, nextTail
	}
}

type Position [2]int

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}
	lines := strings.Split(string(f), "\n")

	head := Position{0, 0}
	tail := Position{0, 0}
	visitedFields := map[string]int{"0:0": 1}

	for _, l := range lines {
		instruction := strings.Fields(l)
		direction := instruction[0]
		steps, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatalf("Error converting steps to integer: %+v", err)
		}

		for i := 0; i < steps; i++ {
			h, t := move(head, tail, direction)
			head = h
			tail = t

			key := fmt.Sprintf("%d:%d", t[0], t[1])
			visitedFields[key]++
		}
	}
	fmt.Printf("Part 1: %d\n", len(visitedFields))
}
