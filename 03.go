package main

func init() {
	addSolutions(3, problem3)
}

func problem3(ctx *problemContext) {
	var banks []batBank
	for line := range ctx.lines() {
		banks = append(banks, batBank(line))
	}
	ctx.reportLoad()

	var part1 int64
	for _, b := range banks {
		part1 += b.maxJoltage(2)
	}
	ctx.reportPart1(part1)

	var part2 int64
	for _, b := range banks {
		part2 += b.maxJoltage(12)
	}
	ctx.reportPart2(part2)
}

type batBank string

func (b batBank) maxJoltage(digits int) int64 {
	var i0 int
	var joltage int64
	for off := range digits {
		d0 := int64(-1)
		var start int
		if off > 0 {
			start = i0 + 1
		}
		end := len(b) - (digits - 1 - off)
		for i := start; i < end; i++ {
			d := int64(b[i] - '0')
			if d > d0 {
				i0 = i
				d0 = d
				if d0 == '9' {
					break
				}
			}
		}
		joltage = 10*joltage + d0
	}
	return joltage
}
