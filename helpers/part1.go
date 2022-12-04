package helpers

import (
	"fmt"
	"os"
)

var cachedPart int = 0

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
