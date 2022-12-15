package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/flopp/aoc2022/helpers"
)

type XY struct {
	x, y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func manhattan(a, b XY) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type Sensor struct {
	pos, beacon XY
}

type Range struct {
	start, end int
}

func computeRanges(sensors []Sensor, y int) ([]Range, int) {
	beacons := make(map[int]bool)
	ranges := make([]Range, 0)
	for _, sensor := range sensors {
		d := manhattan(sensor.pos, sensor.beacon)
		dx := d - abs(y-sensor.pos.y)
		if dx < 0 {
			continue
		}
		if sensor.beacon.y == y {
			beacons[sensor.beacon.x] = true
		}
		ranges = append(ranges, Range{sensor.pos.x - dx, sensor.pos.x + dx})
	}

	sort.Slice(ranges, func(i, j int) bool { return ranges[i].start < ranges[j].start })
	return ranges, len(beacons)
}

func main() {
	reLine := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	sensors := make([]Sensor, 0)
	helpers.ReadStdin(func(line string) {
		if match := reLine.FindStringSubmatch(line); match != nil {
			sensors = append(sensors, Sensor{XY{helpers.MustParseInt(match[1]), helpers.MustParseInt(match[2])}, XY{helpers.MustParseInt(match[3]), helpers.MustParseInt(match[4])}})
		} else {
			panic(fmt.Errorf("bad line: %s", line))
		}
	})

	if helpers.Part1() {
		var y int
		if helpers.Test() {
			y = 10
		} else {
			y = 2_000_000
		}

		ranges, beacons := computeRanges(sensors, y)

		count := 0
		end := 0
		for _, r := range ranges {
			if count == 0 || r.start > end {
				count = 1 + r.end - r.start
				end = r.end
			} else if r.end > end {
				count += r.end - end
				end = r.end
			}
		}

		fmt.Println(count - beacons)
	} else {
		var y1 int
		if helpers.Test() {
			y1 = 20
		} else {
			y1 = 4_000_000
		}

		found := false
		tuningFrequency := 0
		for y := 0; y <= y1; y += 1 {
			ranges, _ := computeRanges(sensors, y)
			first := true
			end := 0
			for _, r := range ranges {
				if first {
					first = false
					end = r.end
				} else if r.end <= end {
					// nothing
				} else if r.start <= end+1 {
					end = r.end
				} else if r.start == end+2 {
					found = true
					tuningFrequency = 4_000_000*(end+1) + y
					break
				}
			}
			if found {
				break
			}
		}

		fmt.Println(tuningFrequency)
	}
}
