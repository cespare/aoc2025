package main

import "github.com/cespare/next/container/set"

func init() {
	addSolutions(4, problem4)
}

func problem4(ctx *problemContext) {
	var r paperRoom
	for line := range ctx.lines() {
		r.g.addRow([]byte(line))
	}
	ctx.reportLoad()

	ctx.reportPart1(r.part1())
	ctx.reportPart2(r.part2())
}

type paperRoom struct {
	g grid[byte]
}

func (r *paperRoom) part1() int {
	var canRemove int
	for v, c := range r.g.all() {
		if c != '@' {
			continue
		}
		var adj int
		for _, n := range v.neighbors8() {
			if !r.g.contains(n) {
				continue
			}
			if r.g.at(n) == '@' {
				adj++
			}
		}
		if adj < 4 {
			canRemove++
		}
	}
	return canRemove
}

func (r *paperRoom) part2() int {
	var removed int
	var q set.Set[vec2]
	for v, c := range r.g.all() {
		if c == '@' {
			q.Add(v)
		}
	}
outer:
	for q.Len() > 0 {
		v := popSet(&q)
		if r.g.at(v) != '@' {
			continue
		}
		neighbors := make([]vec2, 0, 3)
		for _, n := range v.neighbors8() {
			if r.g.contains(n) && r.g.at(n) == '@' {
				if len(neighbors) == 3 {
					continue outer
				}
				neighbors = append(neighbors, n)
			}
		}
		r.g.set(v, 'x')
		removed++
		q.Add(neighbors...)
	}
	return removed
}

func popSet[E comparable](s *set.Set[E]) E {
	for e := range s.All() {
		s.Remove(e)
		return e
	}
	panic("empty")
}
