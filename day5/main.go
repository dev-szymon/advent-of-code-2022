package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

type Stacks [][]string

func (s Stacks) moveCrates(amount, from, to int) {
	index := len(s[from]) - amount
	itemsToMove := s[from][index:]
	s[from] = s[from][:index]
	s[to] = append(s[to], itemsToMove...)
}

func (s Stacks) getTopCrates() string {
	var str string
	for _, stack := range s {
		str += stack[len(stack)-1]
	}
	return str
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	parts := strings.Split(string(f), "\n\n")
	arrangements := strings.Split(parts[0], "\n")
	slices.Reverse(arrangements)

	stacksDescription := arrangements[0]
	stackIds := strings.Fields(stacksDescription)
	stacks := make(Stacks, len(stackIds))

	for i, id := range stackIds {
		stackPosition := strings.Index(stacksDescription, id)
		for _, r := range arrangements[1:] {
			if !unicode.IsSpace(rune(r[stackPosition])) {
				stacks[i] = append(stacks[i], string(r[stackPosition]))
			}
		}
	}

	instructionsLines := strings.Split(parts[1], "\n")
	re := regexp.MustCompile("[0-9]+")

	instructions := make([][]int, len(instructionsLines))
	for i, line := range instructionsLines {
		for _, s := range re.FindAllString(line, 3) {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Error converting instructions. Unexpected value: %s", s)
			}
			instructions[i] = append(instructions[i], n)
		}
	}

	// deep copy slice for part2
	stacks2 := make(Stacks, len(stacks))
	for i, stack := range stacks {
		newStack := make([]string, len(stack))
		copy(newStack, stack)
		stacks2[i] = newStack
	}

	for _, instruction := range instructions {
		amount := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1
		// part1
		for i := 0; i < amount; i++ {
			stacks.moveCrates(1, from, to)
		}
		// part2
		stacks2.moveCrates(amount, from, to)
	}

	fmt.Println(stacks.getTopCrates())
	fmt.Println(stacks2.getTopCrates())
}
