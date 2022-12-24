package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type Blizzard struct {
	dir int
	xy  helpers.XY
}

func (b *Blizzard) advance(width, height int, occupied *[]int) {
	(*occupied)[b.xy.X+b.xy.Y*width] -= 1
	switch b.dir {
	case 0: // N
		b.xy.Y -= 1
		if b.xy.Y == 0 {
			b.xy.Y = height - 2
		}
	case 1: // E
		b.xy.X += 1
		if b.xy.X == width-1 {
			b.xy.X = 1
		}
	case 2: // S
		b.xy.Y += 1
		if b.xy.Y == height-1 {
			b.xy.Y = 1
		}
	case 3: // W
		b.xy.X -= 1
		if b.xy.X == 0 {
			b.xy.X = width - 2
		}
	}
	(*occupied)[b.xy.X+b.xy.Y*width] += 1
}

func main() {
	width := 0
	height := 0
	blizzards := make([]*Blizzard, 0)

	helpers.ReadStdin(func(line string) {
		if strings.HasSuffix(line, "##") {
			// first
			width = len(line)
		} else if strings.HasPrefix(line, "##") {
			// last
		} else {
			for x, c := range line {
				switch c {
				case '^':
					blizzards = append(blizzards, &Blizzard{0, helpers.XY{X: x, Y: height}})
				case '>':
					blizzards = append(blizzards, &Blizzard{1, helpers.XY{X: x, Y: height}})
				case 'v':
					blizzards = append(blizzards, &Blizzard{2, helpers.XY{X: x, Y: height}})
				case '<':
					blizzards = append(blizzards, &Blizzard{3, helpers.XY{X: x, Y: height}})
				}
			}
		}
		height += 1
	})

	start := helpers.XY{X: 1, Y: 0}
	end := helpers.XY{X: width - 2, Y: height - 1}
	occupied := make([]int, width*height)
	for _, b := range blizzards {
		occupied[b.xy.X+b.xy.Y*width] += 1
	}
	for x := 0; x < width; x += 1 {
		if x != start.X {
			occupied[x] = 1
		}
		if x != end.X {
			occupied[x+end.Y*width] = 1
		}
	}
	for y := 0; y < height; y += 1 {
		occupied[y*width] = 1
		occupied[width-1+y*width] = 1
	}

	if helpers.Part1() {
		positions := make(map[helpers.XY]int)
		positions[start] = 0
		for t := 0; true; t += 1 {
			for _, b := range blizzards {
				b.advance(width, height, &occupied)
			}

			next := make(map[helpers.XY]int)
			for p, _ := range positions {
				if p == end {
					fmt.Println(t)
					return
				}
				// stay
				if occupied[p.X+p.Y*width] == 0 {
					next[p] = 0
				}
				// N
				if p.Y > 0 && occupied[p.X+(p.Y-1)*width] == 0 {
					next[p.Offset(0, -1)] = 0
				}
				// E
				if occupied[p.X+1+p.Y*width] == 0 {
					next[p.Offset(+1, 0)] = 0
				}
				// S
				if p.Y < height-1 && occupied[p.X+(p.Y+1)*width] == 0 {
					next[p.Offset(0, +1)] = 0
				}
				// W
				if occupied[p.X-1+p.Y*width] == 0 {
					next[p.Offset(-1, 0)] = 0
				}
			}

			positions = next
		}
	} else {
		type P struct {
			xy  helpers.XY
			run int
		}

		positions := make(map[P]int)
		positions[P{start, 0}] = 0
		for t := 0; true; t += 1 {
			for _, b := range blizzards {
				b.advance(width, height, &occupied)
			}

			next := make(map[P]int)
			for p, _ := range positions {
				if p.run == 0 {
					if p.xy == end {
						p.run = 1
					}
				} else if p.run == 1 {
					if p.xy == start {
						p.run = 2
					}
				} else if p.run == 2 {
					if p.xy == end {
						fmt.Println(t)
						return
					}
				}
				// stay
				if occupied[p.xy.X+p.xy.Y*width] == 0 {
					next[p] = 0
				}
				// N
				if p.xy.Y > 0 && occupied[p.xy.X+(p.xy.Y-1)*width] == 0 {
					next[P{p.xy.Offset(0, -1), p.run}] = 0
				}
				// E
				if occupied[p.xy.X+1+p.xy.Y*width] == 0 {
					next[P{p.xy.Offset(+1, 0), p.run}] = 0
				}
				// S
				if p.xy.Y < height-1 && occupied[p.xy.X+(p.xy.Y+1)*width] == 0 {
					next[P{p.xy.Offset(0, +1), p.run}] = 0
				}
				// W
				if occupied[p.xy.X-1+p.xy.Y*width] == 0 {
					next[P{p.xy.Offset(-1, 0), p.run}] = 0
				}
			}
			positions = next
		}
	}
}
