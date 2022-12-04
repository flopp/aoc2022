package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

func inRange(from, to, candidate int) bool {
	return from <= candidate && candidate <= to
}
func main() {
	lineRxp := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)
	countOverlaps := 0
	helpers.ReadStdin(func(line string) {
		match := lineRxp.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Errorf("bad line: %v", line))
		}
		f0 := helpers.MustParseInt(match[1])
		t0 := helpers.MustParseInt(match[2])
		f1 := helpers.MustParseInt(match[3])
		t1 := helpers.MustParseInt(match[4])

		if helpers.Part1() {
			if (inRange(f0, t0, f1) && inRange(f0, t0, t1)) || (inRange(f1, t1, f0) && inRange(f1, t1, t0)) {
				countOverlaps += 1
			}
		} else {
			if inRange(f0, t0, f1) || inRange(f0, t0, t1) || inRange(f1, t1, f0) || inRange(f1, t1, t0) {
				countOverlaps += 1
			}
		}
	})

	fmt.Println(countOverlaps)
}
