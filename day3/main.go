package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func mustFindCommonRune(args ...string) rune {
	occurances := make(map[rune]map[int]bool)
	for i, s := range args {
		for _, r := range s {
			_, ok := occurances[r]
			if !ok {
				occurances[r] = map[int]bool{i: true}
			} else {
				occurances[r][i] = true
			}
		}
	}
	for r, o := range occurances {
		if len(o) == len(args) {
			return r
		}
	}
	panic("Common rune not found")
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %+v", err)
	}
	items := strings.Split(string(file), "\n")

	var sumPart1 int

	groups := make([][]string, len(items)/3)

	for i, item := range items {
		s1 := item[:len(item)/2]
		s2 := item[len(item)/2:]
		common := mustFindCommonRune(s1, s2)
		sumPart1 += strings.Index(priority, string(common)) + 1

		groups[i/3] = append(groups[i/3], item)
	}
	var sumPart2 int
	for _, g := range groups {
		c := mustFindCommonRune(g...)
		sumPart2 += strings.Index(priority, string(c)) + 1
	}

	fmt.Printf("part1: %d\n", sumPart1)
	fmt.Printf("part2: %d\n", sumPart2)
}
