package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

type Pipes struct {
	start      int
	flows      []int
	tunnels    [][]int
	names      []string
	name2index map[string]int
	candidates []int

	distance [][]int
}

func createPipes() *Pipes {
	return &Pipes{-1, make([]int, 0), make([][]int, 0), make([]string, 0), make(map[string]int), make([]int, 0), make([][]int, 0)}
}

func (pipes *Pipes) index(name string) int {
	i, found := pipes.name2index[name]
	if found {
		return i
	}

	if len(name) != 2 {
		panic(fmt.Errorf("bad valve name '%s'", name))
	}

	i = len(pipes.name2index)
	if name == "AA" {
		pipes.start = i
	}

	pipes.name2index[name] = i
	pipes.names = append(pipes.names, name)
	pipes.flows = append(pipes.flows, 0)
	pipes.tunnels = append(pipes.tunnels, make([]int, 0))

	return i
}

func (pipes *Pipes) addValve(name string, flow int, tunnels []string) {
	i := pipes.index(name)
	pipes.flows[i] = flow
	if flow > 0 {
		pipes.candidates = append(pipes.candidates, i)
	}
	for _, tunnel := range tunnels {
		t := pipes.index(tunnel)
		pipes.tunnels[i] = append(pipes.tunnels[i], t)
	}
}

func (pipes *Pipes) computeDistances() {
	n := len(pipes.tunnels)
	pipes.distance = nil
	for i := 0; i < n; i += 1 {
		d := make([]int, n)
		for j := 0; j < n; j += 1 {
			if i == j {
				d[j] = 0
			} else {
				d[j] = n
			}
		}
		pipes.distance = append(pipes.distance, d)
	}
	for i, tunnels := range pipes.tunnels {
		for _, j := range tunnels {
			pipes.distance[i][j] = 1
		}
	}

	for k := 0; k < n; k += 1 {
		for i := 0; i < n; i += 1 {
			for j := 0; j < n; j += 1 {
				c := pipes.distance[i][k] + pipes.distance[k][j]
				if c < pipes.distance[i][j] {
					pipes.distance[i][j] = c
				}
			}
		}
	}
}

type State struct {
	at    int
	t     int
	open  uint64
	flow  int
	total int
}

func dfs1(pipes *Pipes, timeLimit int, state State, best *int) {
	if state.t >= timeLimit {
		return
	}

	for _, i := range pipes.candidates {
		if (state.open & (1 << i)) != 0 {
			continue
		}

		effort := 1 + pipes.distance[state.at][i]
		if effort <= timeLimit-state.t {
			s := State{i, state.t + effort, state.open | (1 << i), state.flow + pipes.flows[i], state.total + state.flow*effort}
			dfs1(pipes, timeLimit, s, best)
		}
	}

	state.total += (1 + timeLimit - state.t) * state.flow
	if state.total > *best {
		*best = state.total
	}
}

func dfs2(pipes *Pipes, timeLimit int, state State, best *int) {
	if state.t >= timeLimit {
		return
	}

	for _, i := range pipes.candidates {
		if (state.open & (1 << i)) != 0 {
			continue
		}

		effort := 1 + pipes.distance[state.at][i]
		if effort <= timeLimit-state.t {
			s := State{i, state.t + effort, state.open | (1 << i), state.flow + pipes.flows[i], state.total + state.flow*effort}
			dfs2(pipes, timeLimit, s, best)
		}
	}

	state.total += (1 + timeLimit - state.t) * state.flow
	state.at = pipes.start
	state.flow = 0
	state.t = 1
	dfs1(pipes, timeLimit, state, best)
}

func main() {
	pipes := createPipes()
	re := regexp.MustCompile(`^Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)$`)
	helpers.ReadStdin(func(line string) {
		match := re.FindStringSubmatch(line)
		pipes.addValve(match[1], helpers.MustParseInt(match[2]), strings.Split(match[3], ", "))
	})
	pipes.computeDistances()

	open := uint64(0)
	for i, f := range pipes.flows {
		if f == 0 {
			open |= 1 << i
		}
	}
	start := State{pipes.start, 1, open, 0, 0}
	best := 0
	if helpers.Part1() {
		dfs1(pipes, 30, start, &best)
	} else {
		dfs2(pipes, 26, start, &best)
	}
	fmt.Println(best)
}
