package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	partitionIndex := len(s) / 2

	// [inclusive:exclusive]
	a := s[:partitionIndex]
	b := s[partitionIndex:]

	a = mergeSort(a)
	b = mergeSort(b)

	return mergeDescending(a, b)
}

func mergeDescending(a, b []int) []int {
	out := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			out = append(out, a[0])
			a = a[1:]
		} else {
			out = append(out, b[0])
			b = b[1:]
		}
	}
	for len(a) > 0 {
		out = append(out, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		out = append(out, b[0])
		b = b[1:]
	}
	return out
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	var elves []int

	for _, elfCalories := range strings.Split(string(input), "\n\n") {
		var sum int
		for _, m := range strings.Fields(elfCalories) {
			intM, err := strconv.Atoi(m)
			if err != nil {
				log.Fatalf("Error converting meal calories to integer: %+v", intM)
			}
			sum += intM
		}
		elves = append(elves, sum)
	}

	elves = mergeSort(elves)
	var topThreeSum int
	for _, e := range elves[:3] {
		topThreeSum += e
	}
	fmt.Printf("The most calories carried by elf: %d\n", elves[0])
	fmt.Printf("The sum of top three elves' calories: %d\n", topThreeSum)
}
