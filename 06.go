package main

func init() {
	addSolutions(6, problem6)
}

func problem6(ctx *problemContext) {
	g := ctx.byteGrid()
	ctx.reportLoad()

	problems := loadWorksheet(g)

	var part1 int64
	for _, p := range problems {
		part1 += p.computeHoriz()
	}
	ctx.reportPart1(part1)

	var part2 int64
	for _, p := range problems {
		part2 += p.computeVert()
	}
	ctx.reportPart2(part2)
}

type mathProblem struct {
	g  *grid[byte]
	op byte
}

func loadWorksheet(g *grid[byte]) []*mathProblem {
	allSpaces := func(x int64) bool {
		for y := range g.rows {
			if g.at(vec2{x, y}) != ' ' {
				return false
			}
		}
		return true
	}
	var xs []int64
	for x := range g.cols {
		if allSpaces(x) {
			xs = append(xs, x)
		}
	}
	xs = append(xs, g.cols)
	var x0 int64
	var ps []*mathProblem
	for _, x1 := range xs {
		p := &mathProblem{
			g: g.window(
				vec2{x0, 0},
				vec2{x1, g.rows - 1},
			),
		}
		for x := x0; x < x1; x++ {
			if x := g.at(vec2{x, g.rows - 1}); x != ' ' {
				p.op = x
				break
			}
		}
		ps = append(ps, p)
		x0 = x1 + 1
	}
	return ps
}

func (p *mathProblem) computeHoriz() int64 {
	var total int64
	if p.op == '*' {
		total = 1
	}
	for y := range p.g.rows {
		var n int64
		for x := range p.g.cols {
			d := p.g.at(vec2{x, y})
			if d == ' ' {
				continue
			}
			n = n*10 + int64(d-'0')
		}
		if p.op == '+' {
			total += n
		} else {
			total *= n
		}
	}
	return total
}

func (p *mathProblem) computeVert() int64 {
	var total int64
	if p.op == '*' {
		total = 1
	}
	for x := p.g.cols - 1; x >= 0; x-- {
		var n int64
		for y := range p.g.rows {
			d := p.g.at(vec2{x, y})
			if d == ' ' {
				continue
			}
			n = n*10 + int64(d-'0')
		}
		if p.op == '+' {
			total += n
		} else {
			total *= n
		}
	}
	return total
}
