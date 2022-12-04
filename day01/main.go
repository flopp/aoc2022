package main

import (
	"fmt"
	"sort"

	"github.com/flopp/aoc2022/helpers"
)

func main() {
	calories := make([]int, 0)
	current := -1

	helpers.ReadStdin(func(line string) {
		if line == "" {
			current = -1
		} else {
			if current == -1 {
				calories = append(calories, 0)
				current = len(calories) - 1
			}
			calories[current] += helpers.MustParseInt(line)
		}
	})

	sort.Ints(calories)

	last := len(calories) - 1
	if helpers.Part1() {
		fmt.Println(calories[last])
	} else {
		fmt.Println(calories[last] + calories[last-1] + calories[last-2])
	}
}
