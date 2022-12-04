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

func duplicatePriority(contents string) int {
	firstCompartment := make(map[int]int)
	half := len(contents) / 2
	for index, item := range contents {
		p := priority(item)
		if index < half {
			firstCompartment[p] = p
		} else {
			if _, found := firstCompartment[p]; found {
				return p
			}
		}
	}
	return 0
}

func commonPriority(c0, c1, c2 string) int {
	common := make(map[rune]int)

	for _, item := range c0 {
		common[item] = 0
	}
	for _, item := range c1 {
		if _, found := common[item]; found {
			common[item] = 1
		}
	}
	for _, item := range c2 {
		if value, found := common[item]; found && value == 1 {
			return priority(item)
		}
	}

	return 0
}

func main() {
	sum := 0

	if helpers.Part1() {
		helpers.ReadStdin(func(line string) {
			sum += duplicatePriority(line)
		})
	} else {
		count := 0
		group := make([]string, 3)
		helpers.ReadStdin(func(line string) {
			group[count] = line
			count += 1
			if count == 3 {
				sum += commonPriority(group[0], group[1], group[2])
				count = 0
			}
		})
	}

	fmt.Println(sum)
}
