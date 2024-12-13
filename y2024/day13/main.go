package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

//go:embed in.txt
var input string

type Button struct {
	incrX, incrY int
}

type Claw struct {
	buttonA        Button
	buttonB        Button
	prizeX, prizeY int
}

func (claw *Claw) cost(offset int) int {
	ax := claw.buttonA.incrX
	ay := claw.buttonA.incrY
	bx := claw.buttonB.incrX
	by := claw.buttonB.incrY
	px := claw.prizeX + offset
	py := claw.prizeY + offset

	maNum := px*by - py*bx
	maDenom := ax*by - ay*bx
	ma := maNum / maDenom
	if ma*maDenom != maNum {
		return 0
	}

	mbNum := py - ay*ma
	mbDenom := by
	mb := mbNum / mbDenom
	if mb*mbDenom != mbNum {
		return 0
	}

	return 3*ma + mb
}

func parseClaw(input string) *Claw {
	re := regexp.MustCompile("\\d+")
	lines := util.Lines(input)
	aStr, bStr, prizeStr := lines[0], lines[1], lines[2]

	aParts := re.FindAllString(aStr, -1)
	buttonA := Button{
		incrX: conv.ToInt(aParts[0]),
		incrY: conv.ToInt(aParts[1]),
	}

	bParts := re.FindAllString(bStr, -1)
	buttonB := Button{
		incrX: conv.ToInt(bParts[0]),
		incrY: conv.ToInt(bParts[1]),
	}

	prizeParts := re.FindAllString(prizeStr, -1)
	prizeX := conv.ToInt(prizeParts[0])
	prizeY := conv.ToInt(prizeParts[1])

	return &Claw{
		buttonA, buttonB, prizeX, prizeY,
	}
}

func parseInput(input string) []*Claw {
	return arrays.Map(strings.Split(input, "\n\n"), func(s string) *Claw {
		return parseClaw(s)
	})
}

func part1(input string) string {
	return conv.ToString(
		arrays.Sum(
			arrays.Map(parseInput(input), func(c *Claw) int {
				return c.cost(0)
			},
			),
		),
	)
}

func part2(input string) string {
	return conv.ToString(
		arrays.Sum(
			arrays.Map(parseInput(input), func(c *Claw) int {
				return c.cost(10000000000000)
			},
			),
		),
	)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
