package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type XY struct {
	X, Y int
}

func parseXY(s string) XY {
	a := strings.Split(s, ",")
	if len(a) != 2 {
		panic(fmt.Errorf("cannot parse xy: %v", s))
	}
	return XY{helpers.MustParseInt(a[0]), helpers.MustParseInt(a[1])}
}

func (xy XY) offset(x, y int) XY {
	return XY{xy.X + x, xy.Y + y}
}

type Cave struct {
	start XY
	maxY  int
	grid  map[XY]byte
}

func createCave(start XY) *Cave {
	return &Cave{start, start.Y, make(map[XY]byte)}
}

func (cave *Cave) addWall(a, b XY) {
	if a.X == b.X {
		to := helpers.Max(a.Y, b.Y)
		for y := helpers.Min(a.Y, b.Y); y <= to; y += 1 {
			cave.grid[XY{a.X, y}] = '#'
		}
		cave.maxY = helpers.Max(cave.maxY, to)
	} else {
		to := helpers.Max(a.X, b.X)
		for x := helpers.Min(a.X, b.X); x <= to; x += 1 {
			cave.grid[XY{x, a.Y}] = '#'
		}
		cave.maxY = helpers.Max(cave.maxY, a.Y)
	}
}

func (cave *Cave) blocked(xy XY) bool {
	_, found := cave.grid[xy]
	return found
}

// return true if finished
func (cave *Cave) dropSand(part1 bool) bool {
	if !part1 && cave.blocked(cave.start) {
		return true
	}

	sand := cave.start
	for {
		if sand.Y+1 >= cave.maxY+2 {
			if part1 {
				// falls into infinity
				return true
			}
			cave.grid[sand] = 'o'
			return false
		} else if !cave.blocked(sand.offset(0, 1)) {
			sand = sand.offset(0, 1)
		} else if !cave.blocked(sand.offset(-1, 1)) {
			sand = sand.offset(-1, 1)
		} else if !cave.blocked(sand.offset(1, 1)) {
			sand = sand.offset(1, 1)
		} else {
			cave.grid[sand] = 'o'
			return false
		}
	}
}

func main() {
	cave := createCave(XY{500, 0})
	helpers.ReadStdin(func(line string) {
		var last XY
		for i, xys := range strings.Split(line, " -> ") {
			xy := parseXY(xys)
			if i > 0 {
				cave.addWall(last, xy)
			}
			last = xy
		}
	})

	count := 0
	for {
		finished := cave.dropSand(helpers.Part1())
		if finished {
			break
		}
		count += 1
	}
	fmt.Println(count)
}
