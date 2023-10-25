package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Monkey struct {
	id               int
	items            []int
	operands         []string
	arithmeticSign   string
	inspectionCount  int
	isDivisibleValue int
	receiversIds     map[bool]int
}

func (m *Monkey) inspectItem(item int) int {
	m.inspectionCount++
	ops := make([]int, 2)
	for i, o := range m.operands {
		if o == "old" {
			ops[i] = item
		} else {
			n, err := strconv.Atoi(o)
			if err != nil {
				log.Fatalf("Error converting operand to integer: %+v", err)
			}
			ops[i] = n
		}

	}

	if m.arithmeticSign == "+" {
		return ops[0] + ops[1]
	} else if m.arithmeticSign == "*" {
		return ops[0] * ops[1]
	}
	return item
}

func (m *Monkey) testWorryLevel(item int) int {
	isDivisable := item%m.isDivisibleValue == 0
	return m.receiversIds[isDivisable]
}

type Monkeys []*Monkey

func mustFindInt(str string) int {
	re := regexp.MustCompile("[0-9]+")
	s := re.FindString(str)
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("integer not found in string")
	}
	return n
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}
	monkeysMeta := strings.Split(string(f), "\n\n")
	monkeys := make(Monkeys, len(monkeysMeta))

	for _, m := range monkeysMeta {
		description := strings.Split(m, "\n")
		operation := strings.Fields(strings.Split(description[2], "=")[1])
		monkey := &Monkey{
			id:               mustFindInt(description[0]),
			isDivisibleValue: mustFindInt(description[3]),
			arithmeticSign:   operation[1],
			operands:         []string{operation[0], operation[2]},
			receiversIds: map[bool]int{
				true:  mustFindInt(description[4]),
				false: mustFindInt(description[5]),
			},
		}

		items := strings.Split(strings.Split(description[1], ":")[1], ",")
		for _, item := range items {
			n, err := strconv.Atoi(strings.Trim(item, " "))
			if err != nil {
				log.Fatalf("Error converting item worry value to integer: %+v", n)
			}
			monkey.items = append(monkey.items, n)
		}

		monkeys[monkey.id] = monkey
	}

	commonDivisible := 1
	for _, m := range monkeys {
		commonDivisible *= m.isDivisibleValue
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]

			for _, item := range m.items {
				itemWorryLevel := m.inspectItem(item)

				itemWorryLevel %= commonDivisible

				receiverId := m.testWorryLevel(itemWorryLevel)
				monkeys[receiverId].items = append(monkeys[receiverId].items, itemWorryLevel)
				m.items = m.items[1:]
			}
		}
	}

	slices.SortFunc(monkeys, func(a, b *Monkey) int {
		return a.inspectionCount - b.inspectionCount
	})
	monkeyBusines := monkeys[len(monkeys)-1].inspectionCount * monkeys[len(monkeys)-2].inspectionCount

	fmt.Printf("Part2: %d\n", monkeyBusines)
}
