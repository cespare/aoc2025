package main

import "github.com/cespare/next/container/set"

func init() {
	addSolutions(7, problem7)
}

func problem7(ctx *problemContext) {
	g := ctx.byteGrid()
	ctx.reportLoad()

	m := newTachyonManifold(g)

	ctx.reportPart1(m.countSplits())
	ctx.reportPart2(m.countTimelines())
}

type tachyonManifold struct {
	g     *grid[byte]
	start vec2
}

func newTachyonManifold(g *grid[byte]) *tachyonManifold {
	m := &tachyonManifold{g: g}
	for v, c := range g.all() {
		if c == 'S' {
			m.start = v
			break
		}
	}
	return m
}

func (m *tachyonManifold) countSplits() int {
	var numSplits int
	var beams set.Set[vec2]
	var explore func(vec2)
	explore = func(v vec2) {
		if beams.Contains(v) {
			return
		}
		if !m.g.contains(v) {
			return
		}
		switch m.g.at(v) {
		case '.':
			beams.Add(v)
			explore(v.add(south))
		case '^':
			numSplits++
			explore(v.add(west))
			explore(v.add(east))
		default:
			panic("bad")
		}
	}
	explore(m.start.add(south))
	return numSplits
}

func (m *tachyonManifold) countTimelines() int64 {
	memo := make(map[vec2]int64)
	var explore func(vec2) int64
	explore = func(v vec2) int64 {
		n, ok := memo[v]
		if ok {
			return n
		}
		if !m.g.contains(v) {
			return 1
		}
		switch m.g.at(v) {
		case '.':
			n = explore(v.add(south))
		case '^':
			n = explore(v.add(west)) + explore(v.add(east))
		default:
			panic("bad")
		}
		memo[v] = n
		return n
	}
	return explore(m.start.add(south))
}
