package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Circuit struct {
	cycle          int
	register       int
	readOnCycle    int
	signalStrength int
	crt            [][]string
}

const ReadInterval = 40

func (c *Circuit) runCycle() {
	if c.cycle == c.readOnCycle {
		c.signalStrength += c.register * c.cycle
		c.readOnCycle += ReadInterval
	}

	row := (c.cycle - 1) / ReadInterval
	col := (c.cycle - 1) - (row * ReadInterval)

	if c.crt[row] == nil {
		c.crt[row] = make([]string, ReadInterval)
	}
	if c.register-1 == col || c.register == col || c.register+1 == col {
		c.crt[row][col] = "#"
	} else {
		c.crt[row][col] = " "
	}

	c.cycle++
}

func (c *Circuit) registerSignal(v int) {
	c.register += v
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}

	lines := strings.Split(string(f), "\n")

	circuit := &Circuit{
		cycle:          1,
		register:       1,
		readOnCycle:    20,
		signalStrength: 0,
		crt:            make([][]string, 6),
	}

	for i := 0; i < len(lines); i++ {
		circuit.runCycle()

		if lines[i] == "noop" {
			continue
		} else {
			v, err := strconv.Atoi(strings.Fields(lines[i])[1])
			if err != nil {
				log.Fatalf("Error converting string to integer: %+v", err)
			}
			circuit.runCycle()
			circuit.registerSignal(v)
		}
	}

	fmt.Printf("Part1: %d\n", circuit.signalStrength)
	fmt.Println("Part2:")
	for _, y := range circuit.crt {
		fmt.Println(strings.Join(y, ""))
	}
}
