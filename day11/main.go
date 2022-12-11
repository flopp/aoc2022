package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type Monkey struct {
	items               []int64
	operation           byte  // '+', '*'
	operation_value     int64 // 0 for "old"
	divisible_by        int64
	target_true_monkey  int64
	target_false_monkey int64
	inspections         int64
}

func (monkey *Monkey) add_item(item int64) {
	monkey.items = append(monkey.items, item)
}

func (monkey *Monkey) inspect(item int64) int64 {
	monkey.inspections += 1

	value := item
	if monkey.operation_value != 0 {
		value = monkey.operation_value
	}

	switch monkey.operation {
	case '+':
		item += value
	case '*':
		item *= value
	}

	return item
}

func (monkey *Monkey) test_and_throw(item int64, monkeys []*Monkey) {
	if item%monkey.divisible_by == 0 {
		monkeys[monkey.target_true_monkey].add_item(item)
	} else {
		monkeys[monkey.target_false_monkey].add_item(item)
	}
}

func (monkey *Monkey) do_part1(monkeys []*Monkey) {
	items := monkey.items
	monkey.items = monkey.items[:0]

	for _, item := range items {
		item = monkey.inspect(item)
		// "reduce worries"
		item /= 3
		monkey.test_and_throw(item, monkeys)
	}
}

func (monkey *Monkey) do_part2(monkeys []*Monkey, all_divisible_by int64) {
	items := monkey.items
	monkey.items = monkey.items[:0]

	for _, item := range items {
		item = monkey.inspect(item)
		// reduce item by all_divisable_by (modulo artihmetics)
		item %= all_divisible_by
		monkey.test_and_throw(item, monkeys)
	}
}

func monkey_business(monkeys []*Monkey) int64 {
	var inspections []int64
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Slice(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })
	return inspections[0] * inspections[1]
}

func main() {
	reMonkey := regexp.MustCompile(`^Monkey (\d+):$`)
	reItems := regexp.MustCompile(`^\s+Starting items: (.*)$`)
	reOperation := regexp.MustCompile(`^\s+Operation: new = old (\+|\*) (old|\d+)$`)
	reTest := regexp.MustCompile(`^\s+Test: divisible by (\d+)$`)
	reTargetTrue := regexp.MustCompile(`^\s+If true: test_and_throw to monkey (\d+)$`)
	reTargetFalse := regexp.MustCompile(`^\s+If false: test_and_throw to monkey (\d+)$`)
	monkeys := make([]*Monkey, 0)
	monkey := (*Monkey)(nil)
	helpers.ReadStdin(func(line string) {
		if m := reMonkey.FindStringSubmatch(line); m != nil {
			monkey = &Monkey{}
			monkeys = append(monkeys, monkey)
		} else if m := reItems.FindStringSubmatch(line); m != nil {
			for _, item := range strings.Split(m[1], ", ") {
				monkey.items = append(monkey.items, helpers.MustParseInt64(item))
			}
		} else if m := reOperation.FindStringSubmatch(line); m != nil {
			monkey.operation = m[1][0]
			if m[2] == "old" {
				monkey.operation_value = 0
			} else {
				monkey.operation_value = helpers.MustParseInt64(m[2])
			}
		} else if m := reTest.FindStringSubmatch(line); m != nil {
			monkey.divisible_by = helpers.MustParseInt64(m[1])
		} else if m := reTargetTrue.FindStringSubmatch(line); m != nil {
			monkey.target_true_monkey = helpers.MustParseInt64(m[1])
		} else if m := reTargetFalse.FindStringSubmatch(line); m != nil {
			monkey.target_false_monkey = helpers.MustParseInt64(m[1])
		} else if line != "" {
			panic(fmt.Errorf("bad line: %s", line))
		}
	})

	if helpers.Part1() {
		for round := 0; round < 20; round += 1 {
			for _, monkey := range monkeys {
				monkey.do_part1(monkeys)
			}
		}
	} else {
		all_divisible_by := int64(1)
		for _, monkey := range monkeys {
			all_divisible_by *= monkey.divisible_by
		}
		for round := 0; round < 10000; round += 1 {
			for _, monkey := range monkeys {
				monkey.do_part2(monkeys, all_divisible_by)
			}
		}
	}

	fmt.Println(monkey_business(monkeys))
}
