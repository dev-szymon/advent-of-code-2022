package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const (
	Win  = 6
	Draw = 3
	Lose = 0
)
const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

var shapes = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func getRoundPoints(opponentShape int, playerShape int) int {
	var outcome int
	isDifferenceWinning := math.Abs(float64(playerShape)-float64(opponentShape)) == 1
	if opponentShape == playerShape {
		outcome = Draw
	} else if playerShape > opponentShape && isDifferenceWinning || opponentShape > playerShape && !isDifferenceWinning {
		outcome = Win
	} else {
		outcome = Lose
	}

	return playerShape + outcome
}

func mustDecryptShape(opponentShape int, expectedOutcome string) int {
	switch expectedOutcome {
	case "X":
		if opponentShape == 1 {
			return 3
		} else {
			return opponentShape - 1
		}
	case "Y":
		return opponentShape
	case "Z":
		if opponentShape == 3 {
			return 1
		} else {
			return opponentShape + 1
		}
	default:
		panic(fmt.Sprintf("Unexpected input: %d, %s", opponentShape, expectedOutcome))
	}
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	rounds := strings.Split(string(input), "\n")

	// part 1
	var total1 int
	for _, round := range rounds {
		playersShapes := strings.Fields(round)
		points := getRoundPoints(shapes[playersShapes[0]], shapes[playersShapes[1]])
		total1 += points
	}
	fmt.Println(total1)

	// part 2
	var total2 int
	for _, round := range rounds {
		playerShapes := strings.Fields(round)
		decryptedShape := mustDecryptShape(shapes[playerShapes[0]], playerShapes[1])
		total2 += getRoundPoints(shapes[playerShapes[0]], decryptedShape)
	}
	fmt.Println(total2)
}
