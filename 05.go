package main

import (
	"math"
	"strings"

	"rsc.io/omap"
)

func init() {
	addSolutions(5, problem5)
}

func problem5(ctx *problemContext) {
	var intervals [][2]int64
	var ids []int64
	phase := 0
	for line := range ctx.lines() {
		if line == "" {
			phase++
			continue
		}
		switch phase {
		case 0:
			a, b, ok := strings.Cut(line, "-")
			if !ok {
				panic("bad")
			}
			ival := [2]int64{parseInt(a), parseInt(b)}
			intervals = append(intervals, ival)
		case 1:
			ids = append(ids, parseInt(line))
		default:
			panic("bad")
		}
	}
	ctx.reportLoad()

	var s intervalSet
	for _, ival := range intervals {
		s.add([2]int64{ival[0], ival[1] + 1}) // change to half-open
	}

	var part1 int
	for _, id := range ids {
		if s.contains(id) {
			part1++
		}
	}
	ctx.reportPart1(part1)

	ctx.reportPart2(s.count())
}

type intervalSet struct {
	// Non-overlapping, ordered, half-open ranges.
	m omap.Map[int64, int64] // hi -> lo
}

func (s *intervalSet) add(ival [2]int64) {
	var toDelete []int64
	for hi, lo := range s.m.Scan(ival[0], math.MaxInt64) {
		if ival[1] < lo {
			break
		}
		if len(toDelete) == 2 {
			toDelete[1] = hi
		} else {
			toDelete = append(toDelete, hi)
		}
		ival[0] = min(ival[0], lo)
		ival[1] = max(ival[1], hi)
	}
	if len(toDelete) > 0 {
		s.m.DeleteRange(toDelete[0], toDelete[len(toDelete)-1])
	}
	s.m.Set(ival[1], ival[0])
}

func (s *intervalSet) contains(n int64) bool {
	for hi, lo := range s.m.Scan(n, math.MaxInt64) {
		return n >= lo && n < hi
	}
	return false
}

func (s *intervalSet) count() int64 {
	var total int64
	for hi, lo := range s.m.All() {
		total += hi - lo
	}
	return total
}
