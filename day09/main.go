package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

type XY struct {
	X, Y int
}

func (xy *XY) move(direction byte) {
	switch direction {
	case 'U':
		xy.Y -= 1
	case 'D':
		xy.Y += 1
	case 'L':
		xy.X -= 1
	case 'R':
		xy.X += 1
	}
}

func (xy *XY) follow(other XY) {
	dx := other.X - xy.X
	dy := other.Y - xy.Y
	if -1 <= dx && dx <= 1 && -1 <= dy && dy <= 1 {
		return
	}

	if dx == 2 {
		xy.X += 1
	} else if dx == -2 {
		xy.X -= 1
	} else {
		xy.X = other.X
	}

	if dy == 2 {
		xy.Y += 1
	} else if dy == -2 {
		xy.Y -= 1
	} else {
		xy.Y = other.Y
	}
}

type Grid struct {
	Knots         []XY
	VisitedByTail map[XY]bool
}

func createGrid(numberOfKnots int) Grid {
	knots := make([]XY, numberOfKnots)
	grid := Grid{knots, make(map[XY]bool)}
	grid.VisitedByTail[XY{0, 0}] = true
	return grid
}

func (grid *Grid) move(direction byte, distance int) {
	for i := 0; i < distance; i += 1 {
		for j := range grid.Knots {
			if j == 0 {
				grid.Knots[j].move(direction)
			} else {
				grid.Knots[j].follow(grid.Knots[j-1])
			}
		}
		grid.VisitedByTail[grid.Knots[len(grid.Knots)-1]] = true
	}
}

func main() {
	reLine := regexp.MustCompile(`^(U|D|L|R) (\d+)$`)

	knots := 2
	if !helpers.Part1() {
		knots = 10
	}
	grid := createGrid(knots)

	helpers.ReadStdin(func(line string) {
		match := reLine.FindStringSubmatch(line)
		direction := match[1][0]
		distance := helpers.MustParseInt(match[2])
		grid.move(direction, distance)
	})

	fmt.Println(len(grid.VisitedByTail))
}
