package main

func init() {
	addSolutions(1, problem1)
}

func problem1(ctx *problemContext) {
	type dialInst struct {
		right bool
		n     int
	}

	var insts []dialInst
	for line := range ctx.lines() {
		inst := dialInst{
			right: line[0] == 'R',
			n:     int(parseInt(line[1:])),
		}
		insts = append(insts, inst)
	}
	ctx.reportLoad()

	var part1 int
	pos := 50
	for _, inst := range insts {
		if inst.right {
			pos += inst.n
		} else {
			pos -= inst.n
		}
		pos = rem(pos, 100)
		if pos == 0 {
			part1++
		}
	}
	ctx.reportPart1(part1)

	var part2 int
	pos = 50
	for _, inst := range insts {
		if inst.right {
			pos += inst.n
			part2 += pos / 100
			pos = rem(pos, 100)
		} else {
			atZero := pos == 0
			pos -= inst.n
			if pos <= 0 {
				d := pos / -100
				if !atZero {
					d++
				}
				part2 += d
			}
			pos = rem(pos, 100)
		}
	}
	ctx.reportPart2(part2)
}

func rem(n, m int) int {
	r := n % m
	if r < 0 {
		r += m
	}
	return r
}
