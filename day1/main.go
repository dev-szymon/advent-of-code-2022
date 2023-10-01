package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func dayOne() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	var mostCalories int
	elfs := strings.Split(string(input), "\n\n")
	for _, elfCalories := range elfs {
		var sum int
		meals := strings.Split(elfCalories, "\n")
		for _, m := range meals {
			intM, err := strconv.Atoi(m)
			if err != nil {
				log.Fatalf("Error converting meal calorie to integer: %+v", intM)
			}
			sum = sum + intM
		}

		if sum > mostCalories {
			mostCalories = sum
		}
	}

	fmt.Printf("The most calories the elf carries is: %d\n", mostCalories)
}

func main() {
	dayOne()
}
