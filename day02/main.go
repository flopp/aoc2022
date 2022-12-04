package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

func score_part1(opponent byte, you byte) int {
	switch opponent {
	case 'A': // rock
		switch you {
		case 'X':
			return 1 + 3
		case 'Y':
			return 2 + 6
		case 'Z':
			return 3 + 0
		}
	case 'B': // paper
		switch you {
		case 'X':
			return 1 + 0
		case 'Y':
			return 2 + 3
		case 'Z':
			return 3 + 6
		}
	case 'C': // scissors
		switch you {
		case 'X':
			return 1 + 6
		case 'Y':
			return 2 + 0
		case 'Z':
			return 3 + 3
		}
	}

	panic(fmt.Errorf("bad line: %v %v", opponent, you))
}

func score_part2(opponent byte, outcome byte) int {
	switch opponent {
	case 'A': // rock
		switch outcome {
		case 'X': // loose
			return 3 + 0
		case 'Y': // draw
			return 1 + 3
		case 'Z': // win
			return 2 + 6
		}
	case 'B': // paper
		switch outcome {
		case 'X': // loose
			return 1 + 0
		case 'Y': // draw
			return 2 + 3
		case 'Z': // win
			return 3 + 6
		}
	case 'C': // scissors
		switch outcome {
		case 'X': // loose
			return 2 + 0
		case 'Y': // draw
			return 3 + 3
		case 'Z': // win
			return 1 + 6
		}
	}
	panic(fmt.Errorf("bad line: %v %v", opponent, outcome))
}

func main() {
	totalScore := 0
	helpers.ReadStdin(func(line string) {
		if helpers.Part1() {
			totalScore += score_part1(line[0], line[2])
		} else {
			totalScore += score_part2(line[0], line[2])
		}
	})

	fmt.Println(totalScore)
}
