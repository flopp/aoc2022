package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

func pop(s *[]string) string {
	l := len(*s)
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top
}

func push(s *[]string, el string) {
	*s = append(*s, el)
}

func pushEnd(s *[]string, el string) {
	*s = append([]string{el}, (*s)...)
}

func move_part1(stacks [][]string, count, from, to int) {
	for i := 0; i < count; i += 1 {
		el := pop(&(stacks[from-1]))
		push(&(stacks[to-1]), el)
	}
}

func move_part2(stacks [][]string, count, from, to int) {
	t := &(stacks[to-1])
	f := &(stacks[from-1])
	lenf := len(*f)
	*t = append(*t, (*f)[lenf-count:]...)
	*f = (*f)[:lenf-count]
}

func tops(stacks [][]string) string {
	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}

func main() {
	var stacks [][]string

	moveR := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	helpers.ReadStdin(func(line string) {
		if strings.Contains(line, "[") {
			n := (len(line) + 1) / 4
			for len(stacks) < n {
				stacks = append(stacks, nil)
			}
			for i := 0; i < n; i += 1 {
				el := string(line[i*4+1])
				if el != " " {
					pushEnd(&(stacks[i]), el)
				}
			}
			return
		}
		match := moveR.FindStringSubmatch(line)
		if match != nil {
			count := helpers.MustParseInt(match[1])
			from := helpers.MustParseInt(match[2])
			to := helpers.MustParseInt(match[3])
			if helpers.Part1() {
				move_part1(stacks, count, from, to)
			} else {
				move_part2(stacks, count, from, to)
			}
			return
		}
	})

	fmt.Println(tops(stacks))
}
