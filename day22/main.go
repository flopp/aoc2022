package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

type Movement struct {
	turn byte
	dist int
}

type Dir int

const (
	DIR_RIGHT = Dir(0)
	DIR_DOWN  = Dir(1)
	DIR_LEFT  = Dir(2)
	DIR_UP    = Dir(3)
)

func parse(line string) []Movement {
	res := make([]Movement, 0)
	dist := 0
	for i := 0; i < len(line); i += 1 {
		c := line[i]
		if c == 'R' || c == 'L' {
			if dist > 0 {
				res = append(res, Movement{0, dist})
				dist = 0
			}
			res = append(res, Movement{c, 0})
		} else if c >= '0' && c <= '9' {
			dist = dist*10 + int(c-'0')
		} else {
			panic(fmt.Errorf("bad line: %s", line))
		}
	}
	if dist > 0 {
		res = append(res, Movement{0, dist})
	}
	return res
}

func rotate(d Dir, turn byte) Dir {
	i := int(d)
	if turn == 'R' {
		i += 1
	} else {
		i += 3
	}
	return Dir(i % 4)
}

func main() {
	state := 0
	board := make([][]byte, 0)
	width := 0
	movements := make([]Movement, 0)
	helpers.ReadStdin(func(line string) {
		switch state {
		case 0:
			if line == "" {
				state = 1
			} else {
				row := []byte(line)
				width = helpers.Max(width, len(row))
				board = append(board, row)
			}
		case 1:
			movements = parse(line)
			state += 1
		default:
			panic(fmt.Errorf("bad input (state %d): %s", state, line))
		}
	})

	height := len(board)
	// fill rows to width
	for y, row := range board {
		for len(row) < width {
			row = append(row, ' ')
		}
		board[y] = row
	}

	// Find start position
	xy := helpers.XY{X: 0, Y: 0}
	for board[0][xy.X] != '.' {
		xy.X += 1
	}
	d := DIR_RIGHT

	for _, m := range movements {
		if m.dist == 0 {
			d = rotate(d, m.turn)
			board[xy.Y][xy.X] = ">v<^"[int(d)]
		} else {
			switch d {
			case DIR_RIGHT:
				row := board[xy.Y]
				for i := 0; i < m.dist; i += 1 {
					x := xy.X + 1
					for x < width && row[x] == ' ' {
						x += 1
					}
					if x >= width {
						x = 0
					}
					for row[x] == ' ' {
						x += 1
					}
					if row[x] == '#' {
						break
					}
					xy.X = x
				}
			case DIR_LEFT:
				row := board[xy.Y]
				for i := 0; i < m.dist; i += 1 {
					x := xy.X - 1
					for x >= 0 && row[x] == ' ' {
						x -= 1
					}
					if x < 0 {
						x = width - 1
					}
					for row[x] == ' ' {
						x -= 1
					}
					if row[x] == '#' {
						break
					}
					xy.X = x
				}
			case DIR_UP:
				for i := 0; i < m.dist; i += 1 {
					y := xy.Y - 1
					for y >= 0 && board[y][xy.X] == ' ' {
						y -= 1
					}
					if y < 0 {
						y = height - 1
					}
					for board[y][xy.X] == ' ' {
						y -= 1
					}
					if board[y][xy.X] == '#' {
						break
					}
					xy.Y = y
				}
			case DIR_DOWN:
				for i := 0; i < m.dist; i += 1 {
					y := xy.Y + 1
					for y < height && board[y][xy.X] == ' ' {
						y += 1
					}
					if y >= height {
						y = 0
					}
					for board[y][xy.X] == ' ' {
						y += 1
					}
					if board[y][xy.X] == '#' {
						break
					}
					xy.Y = y
				}
			}
		}
	}
	fmt.Println(1000*(xy.Y+1) + 4*(xy.X+1) + int(d))
}
