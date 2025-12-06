package main

import "github.com/cespare/next/container/set"

func init() {
	addSolutions(4, problem4)
}

func problem4(ctx *problemContext) {
	g := ctx.byteGrid()
	ctx.reportLoad()

	ctx.reportPart1(paperRoomPart1(g))
	ctx.reportPart2(paperRoomPart2(g))
}

func paperRoomPart1(g *grid[byte]) int {
	var canRemove int
	for v, c := range g.all() {
		if c != '@' {
			continue
		}
		var adj int
		for _, n := range v.neighbors8() {
			if !g.contains(n) {
				continue
			}
			if g.at(n) == '@' {
				adj++
			}
		}
		if adj < 4 {
			canRemove++
		}
	}
	return canRemove
}

func paperRoomPart2(g *grid[byte]) int {
	var removed int
	var q set.Set[vec2]
	for v, c := range g.all() {
		if c == '@' {
			q.Add(v)
		}
	}
outer:
	for q.Len() > 0 {
		v := popSet(&q)
		if g.at(v) != '@' {
			continue
		}
		neighbors := make([]vec2, 0, 3)
		for _, n := range v.neighbors8() {
			if g.contains(n) && g.at(n) == '@' {
				if len(neighbors) == 3 {
					continue outer
				}
				neighbors = append(neighbors, n)
			}
		}
		g.set(v, 'x')
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
