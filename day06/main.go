package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

func findMarker(s string, size int) int {
	count := make([]int, 26)
	uniques := 0
	lens := len(s)
	for i := 0; i < lens; i += 1 {
		if i-size >= 0 {
			ii := s[i-size] - 'a'
			count[ii] -= 1
			if count[ii] == 0 {
				uniques -= 1
			} else if count[ii] == 1 {
				uniques += 1
			}
		}
		ii := s[i] - 'a'
		count[ii] += 1
		if count[ii] == 1 {
			uniques += 1
		} else if count[ii] == 2 {
			uniques -= 1
		}

		if uniques == size {
			return i + 1
		}
	}
	panic(fmt.Errorf("no 4 unique characters"))
}

func main() {
	helpers.ReadStdin(func(line string) {
		if helpers.Part1() {
			fmt.Println(findMarker(line, 4))
		} else {
			fmt.Println(findMarker(line, 14))
		}
	})
}
