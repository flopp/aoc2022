package main

import (
	"fmt"
	"regexp"

	"github.com/flopp/aoc2022/helpers"
)

type Blueprint struct {
	index                                       int64
	c0min                                       int64
	c00, c01, c02, c03                          int64
	cp2, cp3                                    int64
	max_needed_r0, max_needed_r1, max_needed_r2 int64
}

var (
	bp_c0min                       int64
	bp_c00, bp_c01, bp_c02, bp_c03 int64
	bp_cp2, bp_cp3                 int64
	bp_mr0, bp_mr1, bp_mr2         int64
)

func createBlueprint(i, c00, c01, c02, cp2, c03, cp3 int64) *Blueprint {
	c0 := c00
	if c01 < c0 {
		c0 = c01
	}
	if c02 < c0 {
		c0 = c02
	}
	if c03 < c0 {
		c0 = c03
	}
	mr0 := c00
	if c01 > mr0 {
		mr0 = c01
	}
	if c02 > mr0 {
		mr0 = c02
	}
	if c03 > mr0 {
		mr0 = c03
	}
	mr1 := cp2
	mr2 := cp3

	return &Blueprint{
		i,
		c0,
		c00, c01, c02, c03,
		cp2, cp3,
		mr0, mr1, mr2,
	}
}

func (bp *Blueprint) toGlobal() {
	bp_c0min = bp.c0min
	bp_c00 = bp.c00
	bp_c01 = bp.c01
	bp_c02 = bp.c02
	bp_c03 = bp.c03
	bp_cp2 = bp.cp2
	bp_cp3 = bp.cp3
	bp_mr0 = bp.max_needed_r0
	bp_mr1 = bp.max_needed_r1
	bp_mr2 = bp.max_needed_r2
}

type State struct {
	t              int64
	i0, i1, i2, i3 int64
	r0, r1, r2, r3 int64
}

func createState(timeLimit int64) State {
	s := State{0, 0, 0, 0, 0, 0, 0, 0, 0}
	s.t = timeLimit
	s.r0 += 1
	return s
}

var best int64 = 0

func compute(s State) {
	if s.t == 0 {
		if s.i3 > best {
			best = s.i3
		}
		return
	}

	// we cannot beat the best score even if we build a robot 3 each round
	if best >= s.i3+s.t*s.r3+s.t*(s.t-1)/2 {
		return
	}

	n := s
	n.t -= 1

	// build robot 3 if possible, don't proceed with advancing, or robots 2 + 1 + 0
	if s.i2 >= bp_cp3 && s.i0 >= bp_c03 {
		n.i0 += s.r0 - bp_c03
		n.i1 += s.r1
		n.i2 += s.r2 - bp_cp3
		n.i3 += s.r3
		n.r3 += 1

		compute(n)
		return
	}

	// just advance
	n.i0 += n.r0
	n.i1 += n.r1
	n.i2 += n.r2
	n.i3 += n.r3
	// advance even more if there isn't enough ore (i0) to build any robot
	for n.t > 0 && n.i0 < bp_c0min {
		n.t -= 1
		n.i0 += n.r0
		n.i1 += n.r1
		n.i2 += n.r2
		n.i3 += n.r3
	}
	compute(n)

	// build robot 2 if needed and possible
	if s.r2 < bp_mr2 && s.i1 >= bp_cp2 && s.i0 >= bp_c02 {
		n.i0 = s.i0 + s.r0 - bp_c02
		n.i1 = s.i1 + s.r1 - bp_cp2
		n.i2 = s.i2 + s.r2
		n.i3 = s.i3 + s.r3

		n.r0 = s.r0
		n.r1 = s.r1
		n.r2 = s.r2 + 1
		n.r3 = s.r3
		compute(n)
	}

	// build robot 1 if needed and possible
	if s.r1 < bp_mr1 && s.i0 >= bp_c01 {
		n.i0 = s.i0 + s.r0 - bp_c01
		n.i1 = s.i1 + s.r1
		n.i2 = s.i2 + s.r2
		n.i3 = s.i3 + s.r3

		n.r0 = s.r0
		n.r1 = s.r1 + 1
		n.r2 = s.r2
		n.r3 = s.r3
		compute(n)
	}

	// build robot 0 if needed and possible
	if s.r0 < bp_mr0 && s.i0 >= bp_c00 {
		n.i0 = s.i0 + s.r0 - bp_c00
		n.i1 = s.i1 + s.r1
		n.i2 = s.i2 + s.r2
		n.i3 = s.i3 + s.r3

		n.r0 = s.r0 + 1
		n.r1 = s.r1
		n.r2 = s.r2
		n.r3 = s.r3
		compute(n)
	}
}

func main() {
	reBlueprint := regexp.MustCompile(`^Blueprint (\d+): Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\.$`)
	blueprints := make([]*Blueprint, 0)
	helpers.ReadStdin(func(line string) {
		if match := reBlueprint.FindStringSubmatch(line); match != nil {
			blueprint := createBlueprint(
				helpers.MustParseInt64(match[1]),
				int64(helpers.MustParseInt(match[2])),
				int64(helpers.MustParseInt(match[3])),
				int64(helpers.MustParseInt(match[4])), int64(helpers.MustParseInt(match[5])),
				int64(helpers.MustParseInt(match[6])), int64(helpers.MustParseInt(match[7])),
			)
			blueprints = append(blueprints, blueprint)
		} else {
			panic(fmt.Errorf("bad line: %s", line))
		}
	})

	if helpers.Part1() {
		sum := int64(0)
		for _, bp := range blueprints {
			state := createState(24)
			best = 0
			bp.toGlobal()
			compute(state)
			fmt.Printf("blueprint %d => %d\n", bp.index, best)
			sum += best * bp.index
		}
		fmt.Println(sum)
	} else {
		m := int64(1)

		for i, bp := range blueprints {
			if i >= 3 {
				break
			}
			state := createState(32)
			best = 0
			bp.toGlobal()
			compute(state)
			fmt.Printf("blueprint %d => %d\n", bp.index, best)
			m *= best
		}
		fmt.Println(m)
	}
}
