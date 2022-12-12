package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

type XY struct {
	x, y int
}

func (xy XY) manhattan(other XY) int {
	dx := xy.x - other.x
	if dx < 0 {
		dx = -dx
	}

	dy := xy.y - other.y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

func (xy XY) left() XY {
	return XY{xy.x - 1, xy.y}
}
func (xy XY) right() XY {
	return XY{xy.x + 1, xy.y}
}
func (xy XY) up() XY {
	return XY{xy.x, xy.y - 1}
}
func (xy XY) down() XY {
	return XY{xy.x, xy.y + 1}
}

type Grid struct {
	size   XY
	start  XY
	target XY
	cells  []byte
}

func (grid *Grid) height(xy XY) byte {
	return grid.cells[xy.x+xy.y*grid.size.x]
}

func (grid *Grid) neighbors(xy XY) []XY {
	n := make([]XY, 0, 4)
	h := grid.height(xy)

	if xy.x > 0 {
		o := xy.left()
		if grid.height(o) <= h+1 {
			n = append(n, o)
		}
	}
	if xy.x+1 < grid.size.x {
		o := xy.right()
		if grid.height(o) <= h+1 {
			n = append(n, o)
		}
	}
	if xy.y > 0 {
		o := xy.up()
		if grid.height(o) <= h+1 {
			n = append(n, o)
		}
	}
	if xy.y+1 < grid.size.y {
		o := xy.down()
		if grid.height(o) <= h+1 {
			n = append(n, o)
		}
	}

	return n
}

func reconstruct_path(cameFrom map[XY]XY, current XY) int {
	length := 0
	for {
		prev, found := cameFrom[current]
		if !found {
			break
		}
		length += 1
		current = prev
	}
	return length
}

func (grid *Grid) astar() int {
	openSet := make(map[XY]bool)
	cameFrom := make(map[XY]XY)
	gScore := make(map[XY]int)
	fScore := make(map[XY]int)

	if helpers.Part1() {
		openSet[grid.start] = true
		gScore[grid.start] = 0
		fScore[grid.start] = grid.start.manhattan(grid.target)
	} else {
		for i, h := range grid.cells {
			if h == 'a' {
				xy := XY{i % grid.size.x, i / grid.size.x}
				openSet[xy] = true
				gScore[xy] = 0
				fScore[xy] = xy.manhattan(grid.target)
			}
		}
	}

	for len(openSet) > 0 {
		// can be improved with priority queue
		minF := -1
		var current XY
		for xy := range openSet {
			if f, found := fScore[xy]; found {
				if minF < 0 || f < minF {
					minF = f
					current = xy
				}
			} else {
				panic(fmt.Errorf("inconsistent fScore"))
			}
		}

		if current == grid.target {
			return reconstruct_path(cameFrom, current)
		}

		delete(openSet, current)

		for _, neighbor := range grid.neighbors(current) {
			newG, found := gScore[current]
			if !found {
				panic(fmt.Errorf("inconsistent gScore"))
			}
			newG += 1
			g, found := gScore[neighbor]
			if !found || newG < g {
				cameFrom[neighbor] = current
				gScore[neighbor] = newG
				fScore[neighbor] = newG + neighbor.manhattan(grid.target)
				openSet[neighbor] = true
			}
		}
	}

	return -1
}

func main() {
	grid := Grid{}
	pos := XY{0, 0}
	helpers.ReadStdin(func(line string) {
		for pos.x = 0; pos.x < len(line); pos.x += 1 {
			c := line[pos.x]
			switch c {
			case 'S':
				grid.start = pos
				c = 'a'
			case 'E':
				grid.target = pos
				c = 'z'
			}
			grid.cells = append(grid.cells, c)
		}
		grid.size.x = pos.x
		pos.y += 1
	})
	grid.size.y = pos.y

	fmt.Println(grid.astar())
}
