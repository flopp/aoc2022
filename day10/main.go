package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

func collectSignalStrength(cycle, x int, sum *int) {
	if cycle >= 20 && cycle <= 220 && (cycle-20)%40 == 0 {
		*sum += cycle * x
	}
}

func draw(cycle, x int) {
	cx := (cycle - 1) % 40
	if x-1 <= cx && cx <= x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if cx == 39 {
		fmt.Println("")
	}
}
func main() {
	reAddx := regexp.MustCompile(`^addx (-?\d+)$`)
	reNoop := regexp.MustCompile(`^noop$`)

	x := 1
	cycle := 1

	if helpers.Part1() {
		sum := 0
		helpers.ReadStdin(func(line string) {
			if reNoop.MatchString(line) {
				collectSignalStrength(cycle, x, &sum)
				cycle += 1
			} else if match := reAddx.FindStringSubmatch(line); match != nil {
				collectSignalStrength(cycle, x, &sum)
				cycle += 1
				collectSignalStrength(cycle, x, &sum)
				cycle += 1
				x += helpers.MustParseInt(match[1])
			}
		})
		fmt.Println(sum)
	} else {
		helpers.ReadStdin(func(line string) {
			if reNoop.MatchString(line) {
				draw(cycle, x)
				cycle += 1
			} else if match := reAddx.FindStringSubmatch(line); match != nil {
				draw(cycle, x)
				cycle += 1
				draw(cycle, x)
				cycle += 1
				x += helpers.MustParseInt(match[1])
			}
		})
	}
}
