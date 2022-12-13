package main

import (
	"fmt"
	"sort"

	"github.com/flopp/aoc2022/helpers"
)

type List struct {
	value int
	lists []*List
}

func createNum(i int) *List {
	return &List{i, nil}
}

func createSingletonList(i int) *List {
	list := &List{-1, nil}
	list.append(createNum(i))
	return list
}

func (list *List) append(other *List) {
	list.lists = append(list.lists, other)
}

func pop(stack *[]*List) {
	lenS := len(*stack)
	if lenS == 0 {
		panic(fmt.Errorf("cannot pop from empty stack"))
	}
	*stack = (*stack)[:lenS-1]
}

func push(stack *[]*List, item *List) {
	*stack = append(*stack, item)
}

func top(stack *[]*List) *List {
	lenS := len(*stack)
	if lenS == 0 {
		panic(fmt.Errorf("cannot get top from empty stack"))
	}
	return (*stack)[lenS-1]
}

func parseList(s string) *List {
	root := (*List)(nil)
	stack := make([]*List, 0)

	lenS := len(s)
	for i := 0; i < lenS; /**/ {
		c := s[i]
		if c == '[' {
			list := &List{-1, nil}
			if root == nil {
				root = list
			}
			if len(stack) > 0 {
				top(&stack).append(list)
			}
			push(&stack, list)
			i += 1
		} else if c == ']' {
			pop(&stack)
			i += 1
		} else if '0' <= c && c <= '9' {
			num := ""
			for ; /**/ '0' <= s[i] && s[i] <= '9'; i += 1 {
				num += string(s[i])
			}
			top(&stack).append(createNum(helpers.MustParseInt(num)))
		} else if c == ',' {
			/**/
			i += 1
		} else {
			panic(fmt.Errorf("unsupported char '%s' in '%s'", string(c), s))
		}
	}

	return root
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const (
	INORDER = iota
	NOTINORDER
	UNDECIDED
)

func inOrder(a, b *List) int {
	if a.value >= 0 {
		if b.value >= 0 {
			// value & value
			if a.value < b.value {
				return INORDER
			} else if a.value > b.value {
				return NOTINORDER
			}
			return UNDECIDED
		} else {
			// value & list
			return inOrder(createSingletonList(a.value), b)
		}
	} else if b.value >= 0 {
		// list & value
		return inOrder(a, createSingletonList(b.value))
	} else {
		// list & list
		lenA := len(a.lists)
		lenB := len(b.lists)
		lens := min(len(a.lists), len(b.lists))
		for i := 0; i < lens; i += 1 {
			res := inOrder(a.lists[i], b.lists[i])
			if res != UNDECIDED {
				return res
			}
		}
		if lenA < lenB {
			return INORDER
		} else if lenA > lenB {
			return NOTINORDER
		}
		return UNDECIDED
	}
}

func main() {
	lists := make([]*List, 0)
	helpers.ReadStdin(func(line string) {
		if line != "" {
			lists = append(lists, parseList(line))
		}
	})

	if helpers.Part1() {
		sumInOrder := 0
		for i := 0; i < len(lists)/2; i += 1 {
			if inOrder(lists[2*i], lists[2*i+1]) == INORDER {
				sumInOrder += i + 1
			}
		}
		fmt.Println(sumInOrder)
	} else {
		divider2 := parseList("[[2]]")
		divider6 := parseList("[[6]]")
		lists = append(lists, divider2, divider6)
		sort.Slice(lists, func(i, j int) bool {
			return inOrder(lists[i], lists[j]) == INORDER
		})

		index2 := -1
		index6 := -1
		for i, list := range lists {
			if list == divider2 {
				index2 = i + 1
			} else if list == divider6 {
				index6 = i + 1
			}
		}

		fmt.Println(index2 * index6)
	}
}
