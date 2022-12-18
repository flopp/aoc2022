package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type Cave struct {
	grid []byte
}

func createCave() *Cave {
	return &Cave{make([]byte, 0)}
}

func (cave *Cave) empty() int64 {
	h := len(cave.grid)
	e := int64(0)
	for i := h - 1; i >= 0; i -= 1 {
		if cave.grid[i] != 0 {
			return e
		}
		e += 1
	}
	return e
}

func (cave *Cave) forceEmpty(needed int) {
	count := int64(needed) - cave.empty()
	if count >= 0 {
		for i := int64(0); i < count; i += 1 {
			cave.grid = append(cave.grid, 0)
		}
	} else {
		cave.grid = cave.grid[0 : cave.height()+count]
	}
}

func (cave *Cave) height() int64 {
	return int64(len(cave.grid))
}

func (cave *Cave) collision(rock *Rock, x int, y int64) bool {
	if x < 0 || x+rock.w > 7 || int64(rock.h-1) > y {
		return true
	}

	for _, xy := range rock.g {
		if (cave.grid[y+int64(xy.y)] & (1 << (x + xy.x))) != 0 {
			return true
		}
	}

	return false
}

func (cave *Cave) place(rock *Rock, x int, y int64) {
	for _, xy := range rock.g {
		cave.grid[y+int64(xy.y)] |= (1 << (x + xy.x))
	}
}

func (cave *Cave) print() {
	for i := cave.height() - 1; i >= 0; i -= 1 {
		g := cave.grid[i]
		for x := 0; x < 7; x += 1 {
			if (g & (1 << x)) != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type XY struct {
	x, y int
}
type Rock struct {
	h int
	w int
	g []XY
}

func makeRock(s string) *Rock {
	rows := strings.Split(s, "|")
	h := len(rows)
	w := len(rows[0])
	g := make([]XY, 0)
	for y, row := range rows {
		for x, c := range []byte(row) {
			if c == '#' {
				g = append(g, XY{x, -y})
			}
		}
	}
	return &Rock{h, w, g}
}

func main() {
	winds := ""
	helpers.ReadStdin(func(line string) {
		winds = line
	})

	cave := createCave()

	rocks := make([]*Rock, 5)
	rocks[0] = makeRock("####")
	rocks[1] = makeRock(".#.|###|.#.")
	rocks[2] = makeRock("..#|..#|###")
	rocks[3] = makeRock("#|#|#|#")
	rocks[4] = makeRock("##|##")

	var numberOfRocks int64
	if helpers.Part1() {
		numberOfRocks = 2022
	} else {
		numberOfRocks = 1_000_000_000_000
	}
	windIndex := 0
	for rockIndex := int64(0); rockIndex < numberOfRocks; rockIndex += 1 {
		rockType := rockIndex % int64(len(rocks))
		rock := rocks[rockType]
		cave.forceEmpty(3 + rock.h)

		x := 2
		y := cave.height() - 1
		for {
			wind := winds[windIndex]
			windIndex = (windIndex + 1) % len(winds)

			nx := x
			if wind == '<' {
				nx -= 1
			} else {
				nx += 1
			}

			if !cave.collision(rock, nx, y) {
				x = nx
			}
			if !cave.collision(rock, x, y-1) {
				y -= 1
			} else {
				break
			}
		}
		cave.place(rock, x, y)
		//cave.print()
		//fmt.Println()
	}

	fmt.Println(cave.height() - int64(cave.empty()))
}
