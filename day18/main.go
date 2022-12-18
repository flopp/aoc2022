package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type Range struct {
	min, max int
}

func createRange() Range {
	return Range{1, 0}
}
func (r Range) valid() bool {
	return r.min <= r.max
}
func (r *Range) expand(v int) {
	if !r.valid() {
		r.min = v
		r.max = v
	} else {
		if v < r.min {
			r.min = v
		} else if v > r.max {
			r.max = v
		}
	}
}
func (r Range) contains(v int) bool {
	return r.valid() && r.min <= v && v <= r.max
}

type Cube uint64

func fromXYZ(x, y, z int) Cube {
	if x < -99 || y < -99 || z < -99 || x > 99 || y > 99 || z > 99 {
		panic(fmt.Errorf("bad cube: %d,%d,%d", x, y, z))
	}
	return Cube((uint64(x+127) << 16) | (uint64(y+127) << 8) | (uint64(z+127) << 0))
}

func createCube(line string) Cube {
	a := strings.Split(line, ",")
	x := helpers.MustParseInt(a[0])
	y := helpers.MustParseInt(a[1])
	z := helpers.MustParseInt(a[2])
	return fromXYZ(x, y, z)
}

func (cube Cube) xyz() (int, int, int) {
	x := int((uint64(cube)>>16)&0xFF) - 127
	y := int((uint64(cube)>>8)&0xFF) - 127
	z := int((uint64(cube)>>0)&0xFF) - 127
	return x, y, z
}

func (cube Cube) side(dx, dy, dz int) Cube {
	x, y, z := cube.xyz()
	return fromXYZ(x+dx, y+dy, z+dz)
}

func (cube Cube) in(rx, ry, rz Range) bool {
	x, y, z := cube.xyz()
	return rx.contains(x) && ry.contains(y) && rz.contains(z)
}

func main() {
	cubes := make(map[Cube]bool)
	helpers.ReadStdin(func(line string) {
		cube := createCube(line)
		cubes[cube] = true
	})

	deltas := make([][]int, 6)
	deltas[0] = []int{-1, 0, 0}
	deltas[1] = []int{1, 0, 0}
	deltas[2] = []int{0, -1, 0}
	deltas[3] = []int{0, 1, 0}
	deltas[4] = []int{0, 0, -1}
	deltas[5] = []int{0, 0, 1}

	sides := 0
	if helpers.Part1() {
		for cube := range cubes {
			for _, delta := range deltas {
				cube2 := cube.side(delta[0], delta[1], delta[2])
				if _, found := cubes[cube2]; !found {
					sides += 1
				}
			}
		}
	} else {
		// compute coordinates ranges of all cubes, grow by 1
		rx := createRange()
		ry := createRange()
		rz := createRange()
		for cube := range cubes {
			x, y, z := cube.xyz()
			rx.expand(x)
			ry.expand(y)
			rz.expand(z)
		}
		rx.expand(rx.min - 1)
		rx.expand(rx.max + 1)
		ry.expand(ry.min - 1)
		ry.expand(ry.max + 1)
		rz.expand(rz.min - 1)
		rz.expand(rz.max + 1)

		visited := make(map[Cube]bool)
		pending := make([]Cube, 0)
		corner := fromXYZ(rx.min, ry.min, rz.min)
		pending = append(pending, corner)
		for len(pending) > 0 {
			l := len(pending)
			cube := pending[l-1]
			pending = pending[0 : l-1]

			if _, found := visited[cube]; found {
				continue
			}
			visited[cube] = true

			for _, d := range deltas {
				cube2 := cube.side(d[0], d[1], d[2])
				if cube2.in(rx, ry, rz) {
					if _, found := cubes[cube2]; found {
						sides += 1
					} else {
						pending = append(pending, cube2)
					}
				}
			}
		}
	}
	fmt.Println(sides)
}
