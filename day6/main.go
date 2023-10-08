package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func findIndicator(source string, indicatorLength int) (string, int) {
	var indicator string
	var indicatorEnd int
	for i, r := range source {
		s := string(r)
		if strings.ContainsRune(indicator, r) {
			li := strings.LastIndex(indicator, s)
			indicator = indicator[li+1:] + s
		} else {
			indicator += s
			if len(indicator) == indicatorLength {
				indicatorEnd = i + 1
				break
			}
		}
	}
	return indicator, indicatorEnd
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v\n", err)
	}

	dsb := string(f)
	m1, i1 := findIndicator(dsb, 4)
	m2, i2 := findIndicator(dsb, 14)

	fmt.Printf("part1: %d, %s\n", i1, m1)
	fmt.Printf("part2: %d, %s\n", i2, m2)
}
