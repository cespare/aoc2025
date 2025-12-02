package main

import (
	"strings"
)

func init() {
	addSolutions(2, problem2)
}

func problem2(ctx *problemContext) {
	var ranges [][2]int64
	for chunk := range ctx.splitInput(",") {
		a, b, ok := strings.Cut(chunk, "-")
		if !ok {
			panic(chunk)
		}
		r := [2]int64{parseInt(a), parseInt(b)}
		ranges = append(ranges, r)
	}
	ctx.reportLoad()

	var part1, part2 int64
	for _, r := range ranges {
		for n := r[0]; n <= r[1]; n++ {
			if invalidID1(n) {
				part1 += n
			}
			if invalidID2(n) {
				part2 += n
			}
		}
	}
	ctx.reportPart1(part1)
	ctx.reportPart2(part2)
}

func invalidID1(n int64) bool {
	for m := int64(10); ; m *= 10 {
		a, b := n/m, n%m
		if a == b && a >= m/10 { // check for leading 0
			return true
		}
		if a < b {
			return false
		}
	}
}

func invalidID2(n int64) bool {
	for m := int64(10); ; m *= 10 {
		n, r := n/m, n%m
		if n < r {
			return false
		}
		if r < m/10 { // leading 0
			continue
		}
		for n > 0 {
			if n == r {
				return true
			}
			if n%m != r {
				break
			}
			n /= m
		}
	}
}
