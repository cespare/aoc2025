package main

import (
	"cmp"
	"math"
	"slices"
	"strings"

	"github.com/cespare/next/container/set"
	"rsc.io/top"
)

func init() {
	addSolutions(8, problem8)
}

func problem8(ctx *problemContext) {
	var points []vec3
	for line := range ctx.lines() {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic("bad")
		}
		v := vec3{
			parseInt(parts[0]),
			parseInt(parts[1]),
			parseInt(parts[2]),
		}
		points = append(points, v)
	}
	ctx.reportLoad()

	var (
		pairs       = orderedPairs(points)
		circuits    = make(map[vec3]*set.Set[vec3])
		allCircuits set.Set[*set.Set[vec3]]
		unconnected = len(points)
	)
	for i, p := range pairs {
		v0, v1 := p[0], p[1]
		c0, ok0 := circuits[v0]
		c1, ok1 := circuits[v1]
		switch {
		case !ok0 && !ok1:
			c := set.Of(v0, v1)
			circuits[v0] = c
			circuits[v1] = c
			allCircuits.Add(c)
			unconnected -= 2
		case ok0 && !ok1:
			c0.Add(v1)
			circuits[v1] = c0
			unconnected--
		case !ok0 && ok1:
			c1.Add(v0)
			circuits[v0] = c1
			unconnected--
		case ok0 && ok1:
			if c0 == c1 {
				break
			}
			c := set.Union(c0, c1)
			for v := range c.All() {
				circuits[v] = c
			}
			allCircuits.Remove(c0, c1)
			allCircuits.Add(c)
		default:
			panic("bad")
		}
		if i == 999 {
			top3 := top.New[int](3, cmp.Compare)
			for c := range allCircuits.All() {
				top3.Add(c.Len())
			}
			total := int64(1)
			for _, n := range top3.Take() {
				total *= int64(n)
			}
			ctx.reportPart1(total)
		}
		if allCircuits.Len() == 1 && unconnected == 0 {
			ctx.reportPart2(v0.x * v1.x)
			return
		}
	}
	panic("bad")
}

func orderedPairs(points []vec3) [][2]vec3 {
	var pairs [][2]vec3
	for i, v0 := range points {
		for _, v1 := range points[i+1:] {
			pairs = append(pairs, [2]vec3{v0, v1})
		}
	}
	slices.SortFunc(pairs, func(p0, p1 [2]vec3) int {
		return cmp.Compare(distance3(p0[0], p0[1]), distance3(p1[0], p1[1]))
	})
	return pairs
}

func distance3(v0, v1 vec3) float64 {
	dx := float64(v0.x - v1.x)
	dy := float64(v0.y - v1.y)
	dz := float64(v0.z - v1.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
