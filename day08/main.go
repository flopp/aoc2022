package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

func isVisible(grid []string, x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}
	height := len(grid)
	width := len(grid[0])
	if x+1 == width || y+1 == height {
		return true
	}

	t := grid[y][x]

	allLower := true
	for xx := 0; xx < x; xx += 1 {
		if grid[y][xx] >= t {
			allLower = false
			break
		}
	}
	if allLower {
		return true
	}

	allLower = true
	for xx := x + 1; xx < width; xx += 1 {
		if grid[y][xx] >= t {
			allLower = false
			break
		}
	}
	if allLower {
		return true
	}

	allLower = true
	for yy := 0; yy < y; yy += 1 {
		if grid[yy][x] >= t {
			allLower = false
			break
		}
	}
	if allLower {
		return true
	}

	allLower = true
	for yy := y + 1; yy < height; yy += 1 {
		if grid[yy][x] >= t {
			allLower = false
			break
		}
	}
	if allLower {
		return true
	}

	return false
}

func scenicScore(grid []string, x, y int) int {
	result := 1

	height := len(grid)
	width := len(grid[0])
	t := grid[y][x]

	// left
	s := 0
	for xx := x - 1; xx >= 0; xx -= 1 {
		s += 1
		if grid[y][xx] >= t {
			break
		}
	}
	if s > 0 {
		result *= s
	}

	// right
	s = 0
	for xx := x + 1; xx < width; xx += 1 {
		s += 1
		if grid[y][xx] >= t {
			break
		}
	}
	if s > 0 {
		result *= s
	}

	// up
	s = 0
	for yy := y - 1; yy >= 0; yy -= 1 {
		s += 1
		if grid[yy][x] >= t {
			break
		}
	}
	if s > 0 {
		result *= s
	}

	// down
	s = 0
	for yy := y + 1; yy < height; yy += 1 {
		s += 1
		if grid[yy][x] >= t {
			break
		}
	}
	if s > 0 {
		result *= s
	}

	return result
}

func main() {
	grid := make([]string, 0)
	helpers.ReadStdin(func(line string) {
		grid = append(grid, line)
	})

	height := len(grid)
	width := len(grid[0])

	if helpers.Part1() {
		visible := 0
		for y := 0; y < height; y += 1 {
			for x := 0; x < width; x += 1 {
				if isVisible(grid, x, y) {
					visible += 1
				}
			}
		}
		fmt.Println(visible)
	} else {
		maxScenicScore := 0
		for y := 0; y < height; y += 1 {
			for x := 0; x < width; x += 1 {
				maxScenicScore = helpers.Max(maxScenicScore, scenicScore(grid, x, y))
			}
		}
		fmt.Println(maxScenicScore)
	}
}
