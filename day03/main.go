package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return (int)(item-'a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return (int)(item-'A') + 27
	}
	panic(fmt.Errorf("bad item: %v", item))
}

func commonRune(strings []string) rune {
	seen := make(map[rune]int)
	for index, s := range strings {
		if index == 0 {
			for _, r := range s {
				seen[r] = index
			}
		} else if index+1 < len(strings) {
			for _, r := range s {
				if lastIndex, found := seen[r]; found && lastIndex+1 == index {
					seen[r] = index
				}
			}
		} else {
			for _, r := range s {
				if lastIndex, found := seen[r]; found && lastIndex+1 == index {
					return r
				}
			}
		}
	}
	panic(fmt.Errorf("no common runes"))
}

func main() {
	sum := 0

	if helpers.Part1() {
		helpers.ReadStdin(func(line string) {
			sum += priority(commonRune([]string{line[:len(line)/2], line[len(line)/2:]}))
		})
	} else {
		group := make([]string, 0)
		helpers.ReadStdin(func(line string) {
			group = append(group, line)
			if len(group) == 3 {
				sum += priority(commonRune(group))
				group = group[:0]
			}
		})
	}

	fmt.Println(sum)
}
