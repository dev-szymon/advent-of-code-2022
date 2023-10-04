package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isFullyContained(pair [][]int) bool {
	firstStart := pair[0][0]
	firstEnd := pair[0][1]
	secondStart := pair[1][0]
	secondEnd := pair[1][1]

	return (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd)
}
func isOverlapping(pair [][]int) bool {
	firstStart := pair[0][0]
	firstEnd := pair[0][1]
	secondStart := pair[1][0]
	secondEnd := pair[1][1]

	return (firstStart >= secondStart && firstStart <= secondEnd) || (secondStart >= firstStart && secondStart <= firstEnd)
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	pairs := strings.Split(string(f), "\n")

	var containedPairs int
	var overlappingPairs int
	for _, p := range pairs {
		elves := strings.Split(p, ",")
		pairBoundries := make([][]int, len(elves))
		for i, e := range elves {
			elfBoundries := strings.Split(e, "-")
			for _, b := range elfBoundries {
				boundry, err := strconv.Atoi(b)
				if err != nil {
					log.Fatalf("Unexpected section id: %s", b)
				}
				pairBoundries[i] = append(pairBoundries[i], boundry)
			}

		}

		if isFullyContained(pairBoundries) {
			containedPairs++
		}
		if isOverlapping(pairBoundries) {
			overlappingPairs++
		}
	}

	fmt.Printf("Contained: %d\n", containedPairs)
	fmt.Printf("Overlapping: %d\n", overlappingPairs)
}
