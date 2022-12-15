package helpers

import (
	"fmt"
	"os"
)

var cachedPart int = 0
var cachedTest int = 0

func Part1() bool {
	if cachedPart != 0 {
		return cachedPart == 1
	}

	switch {
	case os.Args[1] == "part1":
		cachedPart = 1
		return true
	case os.Args[1] == "part2":
		cachedPart = 2
		return false
	default:
		panic(fmt.Errorf("bad part option: <%s>", os.Args[1]))
	}
}

func Test() bool {
	if cachedTest != 0 {
		return cachedTest == 1
	}

	if len(os.Args) != 3 {
		panic(fmt.Errorf("bad command line: %v", os.Args))
	}

	switch {
	case os.Args[2] == "test":
		cachedTest = 1
		return true
	case os.Args[2] == "puzzle":
		cachedTest = 2
		return false
	default:
		panic(fmt.Errorf("bad test/puzzle option: <%s>", os.Args[2]))
	}
}

func Puzzle() bool {
	return !Test()
}
