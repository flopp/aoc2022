package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

type Elf struct {
	p helpers.XY
	n helpers.XY
}

func free(occupied map[helpers.XY]bool, a helpers.XY, b helpers.XY, c helpers.XY) bool {
	if _, found := occupied[a]; found {
		return false
	}
	if _, found := occupied[b]; found {
		return false
	}
	if _, found := occupied[c]; found {
		return false
	}
	return true
}

func (elf *Elf) dontMove() {
	elf.n = elf.p
}

func (elf *Elf) determineNextMove(occupied map[helpers.XY]bool, startDir int) {
	candidates := make([]helpers.XY, 0, 4)
	for d := 0; d < 4; d += 1 {
		dir := (startDir + d) % 4
		switch dir {
		case 0: // N
			n := elf.p.Offset(0, -1)
			if free(occupied, elf.p.Offset(-1, -1), n, elf.p.Offset(+1, -1)) {
				candidates = append(candidates, n)
			}
		case 1: // S
			n := elf.p.Offset(0, +1)
			if free(occupied, elf.p.Offset(-1, +1), n, elf.p.Offset(+1, +1)) {
				candidates = append(candidates, n)
			}
		case 2: // W
			n := elf.p.Offset(-1, 0)
			if free(occupied, elf.p.Offset(-1, -1), n, elf.p.Offset(-1, +1)) {
				candidates = append(candidates, n)
			}
		case 3: // E
			n := elf.p.Offset(+1, 0)
			if free(occupied, elf.p.Offset(+1, -1), n, elf.p.Offset(+1, +1)) {
				candidates = append(candidates, n)
			}
		}
	}

	if len(candidates) == 0 || len(candidates) == 4 {
		elf.dontMove()
	} else {
		elf.n = candidates[0]
	}
}

type Range struct {
	from, to int
}

func emptyRange() Range {
	return Range{1, 0}
}
func (r *Range) extend(v int) {
	if r.from > r.to {
		r.from = v
		r.to = v
	}
	r.from = helpers.Min(r.from, v)
	r.to = helpers.Max(r.to, v)
}

func main() {
	elves := make([]*Elf, 0)
	y := 0
	helpers.ReadStdin(func(line string) {
		for x, c := range line {
			if c == '#' {
				elves = append(elves, &Elf{helpers.XY{X: x, Y: y}, helpers.XY{X: x, Y: y}})
			}
		}
		y += 1
	})

	dir := 0

	if helpers.Part1() {
		for round := 0; round < 10; round += 1 {
			// get current positions
			current := make(map[helpers.XY]bool)
			for _, elf := range elves {
				current[elf.p] = true
			}

			// determine next positions
			next := make(map[helpers.XY]int)
			for _, elf := range elves {
				elf.determineNextMove(current, dir)
				next[elf.n] += 1
			}

			// execute moves
			for _, elf := range elves {
				if count := next[elf.n]; count == 1 {
					elf.p = elf.n
				}
			}

			dir = (dir + 1) % 4
		}

		rx := emptyRange()
		ry := emptyRange()
		for _, elf := range elves {
			rx.extend(elf.p.X)
			ry.extend(elf.p.Y)
		}
		fmt.Println((1+rx.to-rx.from)*(1+ry.to-ry.from) - len(elves))
	} else {
		for round := 0; true; round += 1 {
			// get current positions
			current := make(map[helpers.XY]bool)
			for _, elf := range elves {
				current[elf.p] = true
			}

			// determine next positions
			next := make(map[helpers.XY]int)
			for _, elf := range elves {
				elf.determineNextMove(current, dir)
				next[elf.n] += 1
			}

			// execute moves
			moved := false
			for _, elf := range elves {
				if count := next[elf.n]; count == 1 {
					if !moved && elf.p != elf.n {
						moved = true
					}
					elf.p = elf.n
				}
			}
			if !moved {
				fmt.Println(round + 1)
				break
			}

			dir = (dir + 1) % 4
		}
	}
}
