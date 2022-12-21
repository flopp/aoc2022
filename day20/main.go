package main

import (
	"fmt"

	"github.com/flopp/aoc2022/helpers"
)

type Node struct {
	number int
	moves  int
	p, n   *Node
}

func main() {
	rounds := 1
	factor := 1
	if !helpers.Part1() {
		factor = 811589153
		rounds = 10
	}

	start := (*Node)(nil)
	nodes := make([]*Node, 0)
	helpers.ReadStdin(func(line string) {
		number := factor * helpers.MustParseInt(line)
		node := &Node{number, 0, nil, nil}
		if start == nil {
			node.p = node
			node.n = node
			start = node
		} else {
			node.p = start.p
			node.n = start
			start.p.n = node
			start.p = node
		}
		nodes = append(nodes, node)
	})

	// optimize node moves (= number mod count-1)
	count := len(nodes)
	for _, node := range nodes {
		node.moves = node.number % (count - 1)
		if node.moves > count/2 {
			node.moves -= count - 1
		} else if node.moves < -count/2 {
			node.moves += count - 1
		}
	}

	for round := 0; round < rounds; round += 1 {
		for _, node := range nodes {
			moves := node.moves

			for moves > 0 {
				p := node.p
				n := node.n
				nn := n.n

				p.n = n
				n.p = p
				n.n = node
				node.p = n
				node.n = nn
				nn.p = node
				moves -= 1
			}
			for moves < 0 {
				p := node.p
				pp := p.p
				n := node.n

				pp.n = node
				node.p = pp
				node.n = p
				p.p = node
				n.p = p
				p.n = n
				moves += 1
			}
		}
	}

	// find 0 node
	node := start
	for node.number != 0 {
		node = node.n
	}
	// compute sum
	sum := 0
	for n := 0; n < 3000; n += 1 {
		node = node.n
		if ((n + 1) % 1000) == 0 {
			sum += node.number
		}
	}
	fmt.Println(sum)
}
