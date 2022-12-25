package main

import (
	"fmt"
	"strings"

	"github.com/flopp/aoc2022/helpers"
)

func snafu2int(s string) int {
	v := 0
	for _, c := range []byte(s) {
		v = v*5 + strings.IndexByte("=-012", c) - 2
	}
	return v
}

func int2snafu(i int) string {
	if i <= 0 {
		return "0"
	}
	digits := make([]byte, 0)
	for i > 0 {
		d := i % 5
		if d <= 2 {
			digits = append(digits, "012"[d])
		} else {
			digits = append(digits, "=-"[d-3])
			i += 5
		}
		i /= 5
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

func main() {
	sum := 0
	helpers.ReadStdin(func(line string) {
		sum += snafu2int(line)
	})
	fmt.Println(int2snafu(sum))
}
