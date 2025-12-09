package main

import (
	"slices"
	"strings"

	"github.com/cespare/next/container/set"
)

func init() {
	addSolutions(9, problem9)
}

func problem9(ctx *problemContext) {
	var red []vec2
	for line := range ctx.lines() {
		a, b, ok := strings.Cut(line, ",")
		if !ok {
			panic("bad")
		}
		v := vec2{parseInt(a), parseInt(b)}
		red = append(red, v)
	}
	ctx.reportLoad()

	ctx.reportPart1(maxFloorRect(red))
	ctx.reportPart2(maxRedGreen(red))
}

func maxFloorRect(points []vec2) int64 {
	var best int64
	for i, p0 := range points {
		for _, p1 := range points[i+1:] {
			area := (abs(p1.x-p0.x) + 1) * (abs(p1.y-p0.y) + 1)
			best = max(best, area)
		}
	}
	return best
}

func maxRedGreen(points []vec2) int64 {
	// Virtual coordinates: map real x/y to virtual x/y (0, 1, ...).
	var allx, ally set.Set[int64]
	for _, p := range points {
		allx.Add(p.x)
		ally.Add(p.y)
	}
	assignVirt := func(allz *set.Set[int64]) map[int64]int64 {
		zs := slices.Sorted(allz.All())
		var vz int64
		virt := make(map[int64]int64)
		for i, z := range zs {
			if i > 0 {
				d := z - zs[i-1]
				if d == 1 {
					vz++
				} else {
					vz += 2 // collapse >2 into 2
				}
			}
			virt[z] = vz
		}
		return virt
	}
	virtx := assignVirt(&allx)
	virty := assignVirt(&ally)
	toVirt := func(v vec2) vec2 {
		return vec2{virtx[v.x], virty[v.y]}
	}

	// Compute the result in virtual coordinates.
	var bounds rect
	var colored set.Set[vec2]
	prev := toVirt(points[len(points)-1])
	for i, p := range points {
		v := toVirt(p)

		if i == 0 {
			bounds = rect{v, v.add(vec2{1, 1})}
		} else {
			bounds.v0 = bounds.v0.min(v)
			bounds.v1 = bounds.v1.max(v.add(vec2{1, 1}))
		}

		d := v.sub(prev)
		if d.x != 0 {
			d.x /= abs(d.x)
		}
		if d.y != 0 {
			d.y /= abs(d.y)
		}
		for v1 := prev; v1 != v; v1 = v1.add(d) {
			colored.Add(v1)
		}

		prev = v
	}

	// Figure out where the interior vs. exterior is.
	flood := func(v vec2) (area *set.Set[vec2], inside bool) {
		inside = true
		q := []vec2{v}
		area = set.Of(v)
		for len(q) > 0 {
			v := SlicePop(&q)
			for _, n := range v.neighbors4() {
				if colored.Contains(n) || area.Contains(n) {
					continue
				}
				if !bounds.contains(n) {
					inside = false
					continue
				}
				area.Add(n)
				q = append(q, n)
			}
		}
		return area, inside
	}
	var exterior set.Set[vec2]
	for v := range bounds.all() {
		if colored.Contains(v) || exterior.Contains(v) {
			continue
		}
		area, inside := flood(v)
		if inside {
			colored.AddSet(area)
			break
		}
		exterior.AddSet(area)
	}

	// Now compute the answer but check validity of each rect in virtual
	// coordinates.
	var best int64
	for i, p0 := range points {
		for _, p1 := range points[i+1:] {
			r := makeRect(toVirt(p0), toVirt(p1))
			ok := true
			for v := range r.all() {
				if !colored.Contains(v) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			area := (abs(p1.x-p0.x) + 1) * (abs(p1.y-p0.y) + 1)
			best = max(best, area)
		}
	}

	return best
}
