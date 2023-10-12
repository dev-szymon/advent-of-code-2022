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

type Position [2]int

func (p Position) getAbsDistTo(to Position) (int, int) {
	var diffY int
	var diffX int
	hy := float64(p[0])
	hx := float64(p[1])
	ty := float64(to[0])
	tx := float64(to[1])

	if p[0] < to[0] {
		diffY = int(math.Abs(ty - hy))
	} else {
		diffY = int(math.Abs(hy - ty))
	}
	if p[1] < to[1] {
		diffX = int(math.Abs(tx - hx))
	} else {
		diffX = int(math.Abs(hx - tx))
	}
	return diffY, diffX
}

func (p Position) getOffsetTo(to Position) (int, int) {
	y, x := p.getAbsDistTo(to)
	if p[0] < to[0] {
		y = y * -1
	}
	if p[1] < to[1] {
		x = x * -1
	}
	return y, x
}

type Knot struct {
	id           int
	next         *Knot
	prev         *Knot
	position     Position
	prevPosition Position
}

func (r *Knot) move(direction string) {
	y, x := getDirection(direction)
	var nextPosition Position
	nextPosition[0] = r.position[0] + y
	nextPosition[1] = r.position[1] + x
	r.prevPosition = r.position
	r.position = nextPosition
}

func (r *Knot) follow() {
	prevKnotMoveY, prevKnotMoveX := r.prev.position.getOffsetTo(r.prev.prevPosition)
	hasPrevKnotMovedDiagonaly := prevKnotMoveY != 0 && prevKnotMoveX != 0

	offsetY, offsetX := r.prev.position.getOffsetTo(r.position)
	distY, distX := r.position.getAbsDistTo(r.prev.position)
	isAdjacent := distY < 2 && distX < 2

	willBeSameRow := offsetY == 0
	willBeSameCol := offsetX == 0

	var nextPosition Position

	if isAdjacent {
		return
	} else if hasPrevKnotMovedDiagonaly && willBeSameCol {
		nextPosition[1] = r.position[1]
		if r.prev.position[0] > r.position[0] {
			nextPosition[0] = r.prev.position[0] - 1
		} else {
			nextPosition[0] = r.prev.position[0] + 1
		}
	} else if hasPrevKnotMovedDiagonaly && willBeSameRow {
		nextPosition[0] = r.position[0]
		if r.prev.position[1] > r.position[1] {
			nextPosition[1] = r.prev.position[1] - 1
		} else {
			nextPosition[1] = r.prev.position[1] + 1
		}
	} else if hasPrevKnotMovedDiagonaly {
		nextPosition[0] = r.position[0] + prevKnotMoveY
		nextPosition[1] = r.position[1] + prevKnotMoveX
	} else {
		nextPosition[0] = r.prev.prevPosition[0]
		nextPosition[1] = r.prev.prevPosition[1]
	}
	r.prevPosition = r.position
	r.position = nextPosition
}

const (
	initialY = 0
	initialX = 0
)

func newKnot(id int, p Position, prev, next *Knot) *Knot {
	return &Knot{
		id:           id,
		next:         next,
		prev:         prev,
		position:     p,
		prevPosition: p,
	}
}

func prepare(count int, startingPosition Position) *Knot {
	head := newKnot(0, startingPosition, nil, nil)
	curr := head
	for i := 0; i < count-1; i++ {
		nextKnot := newKnot(i+1, startingPosition, curr, nil)
		curr.next = nextKnot
		curr = nextKnot
	}
	return head
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %+v", err)
	}
	lines := strings.Split(string(f), "\n")

	head1 := prepare(2, Position{15, 11})
	head2 := prepare(10, Position{15, 11})

	tail1visited := make(map[Position]bool)
	tail2visited := make(map[Position]bool)

	for _, l := range lines {
		instruction := strings.Fields(l)
		direction := instruction[0]
		steps, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatalf("Error converting steps to integer: %+v", err)
		}

		for m := 0; m < steps; m++ {
			head1.move(direction)
			curr := head1
			for curr.next != nil {
				curr = curr.next
				curr.follow()
				if curr.next == nil {
					tail1visited[curr.position] = true
				}
			}

			head2.move(direction)
			curr2 := head2

			for curr2.next != nil {
				curr2 = curr2.next
				curr2.follow()
				if curr2.next == nil {
					tail2visited[curr2.position] = true
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(tail1visited))
	fmt.Printf("Part 2: %d\n", len(tail2visited))
}

// Utility that helps visualise the current state of the knots
// w 26 h 22 for pt2
func PrintCurrentState(head *Knot, width, height int) {
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		row := make([]string, width)
		for j := range row {
			row[j] = "."
		}
		grid[i] = row
	}
	grid[initialY][initialX] = "s"

	curr := head
	for curr != nil {
		y := curr.position[0]
		x := curr.position[1]
		if grid[y][x] == "." || grid[y][x] == "s" {
			if curr.id == 0 {
				grid[y][x] = "H"
			} else if curr.id == 9 {
				grid[y][x] = "t"
			} else {
				grid[y][x] = fmt.Sprintf("%d", curr.id)
			}
		}
		curr = curr.next
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == width-1 {
				fmt.Printf("%s\n", grid[y][x])
			} else {
				fmt.Printf("%s", grid[y][x])
			}
		}
	}
}
