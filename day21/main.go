package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

type Node struct {
	name string
	leaf bool

	number int64

	name1     string
	operation byte
	name2     string
}

func main() {
	reNumber := regexp.MustCompile(`^(....): (-?\d+)$`)
	reOperation := regexp.MustCompile(`^(....): (....) (.) (....)$`)

	nodes := make(map[string]*Node)
	helpers.ReadStdin(func(line string) {
		if match := reNumber.FindStringSubmatch(line); match != nil {
			name := match[1]
			number := helpers.MustParseInt64(match[2])
			nodes[name] = &Node{name, true, number, "", 0, ""}
		} else if match := reOperation.FindStringSubmatch(line); match != nil {
			name := match[1]
			name1 := match[2]
			operation := match[3][0]
			name2 := match[4]
			nodes[name] = &Node{name, false, 0, name1, operation, name2}
		} else {
			panic(fmt.Errorf("bad line: %s", line))
		}
	})

	topsort := make([]*Node, 0, len(nodes))
	sorted := make(map[string]int)
	pending := make([]*Node, 0, len(nodes))
	for name, node := range nodes {
		if node.leaf {
			sorted[name] = len(topsort)
			topsort = append(topsort, node)
		} else {
			pending = append(pending, node)
		}
	}

	pending2 := make([]*Node, 0, len(pending))
	for len(pending) > 0 {
		pending2 = pending2[:0]
		for _, node := range pending {
			if _, found := sorted[node.name1]; !found {
				pending2 = append(pending2, node)
			} else if _, found := sorted[node.name2]; !found {
				pending2 = append(pending2, node)
			} else {
				sorted[node.name] = len(topsort)
				topsort = append(topsort, node)
			}
		}
		pending = pending2
	}

	value := make([]int64, len(topsort))
	type Operation struct {
		op1, op2 int
		op       byte
	}
	offset := 0
	// TODO: detect and prune unneeded operations (that are not in the "cone" of root)
	operations := make([]Operation, 0)
	for i, node := range topsort {
		if node.leaf {
			offset += 1
			value[i] = node.number
		} else {
			operations = append(operations, Operation{sorted[node.name1], sorted[node.name2], node.operation})
		}
	}

	if helpers.Part1() {
		for i, operation := range operations {
			switch operation.op {
			case '+':
				value[offset+i] = value[operation.op1] + value[operation.op2]
			case '-':
				value[offset+i] = value[operation.op1] - value[operation.op2]
			case '*':
				value[offset+i] = value[operation.op1] * value[operation.op2]
			case '/':
				value[offset+i] = value[operation.op1] / value[operation.op2]
			}
		}
		fmt.Println(value[len(value)-1])
	} else {
		humni := sorted["humn"]
		root_op1 := operations[len(operations)-1].op1
		root_op2 := operations[len(operations)-1].op2
		operations = operations[0 : len(operations)-1]

		// manual search using increasing start vaklues & decreaing increments
		for humn := int64(3429411069020); true; humn += 1 {
			value[humni] = humn
			for i, operation := range operations {
				v1 := value[operation.op1]
				v2 := value[operation.op2]
				switch operation.op {
				case '+':
					value[offset+i] = v1 + v2
				case '-':
					value[offset+i] = v1 - v2
				case '*':
					value[offset+i] = v1 * v2
				case '/':
					value[offset+i] = v1 / v2
				}
			}
			v1 := value[root_op1]
			v2 := value[root_op2]
			fmt.Printf("%v => %v %v\n", humn, v1, v2)
			if v1 == v2 {
				fmt.Println(humn)
				break
			}
			if v1 < v2 {
				break
			}
		}
	}

}
