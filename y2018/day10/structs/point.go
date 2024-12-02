package structs

import (
	"adventOfCode/util"
	"adventOfCode/util/conv"
	"strings"
)

type Point struct {
	X, Y, Vx, Vy int
}

func NewPoint(line string) *Point {
	line = strings.ReplaceAll(line, "position=<", "")
	line = strings.ReplaceAll(line, "> velocity=<", " ")
	line = strings.ReplaceAll(line, ",", " ")
	line = strings.ReplaceAll(line, ">", "")
	line = strings.TrimSpace(line)

	parts := util.SplitSpace(line)
	return &Point{
		X:  conv.ToInt(parts[0]),
		Y:  conv.ToInt(parts[1]),
		Vx: conv.ToInt(parts[2]),
		Vy: conv.ToInt(parts[3]),
	}
}

func (p *Point) Update() {
	p.X += p.Vx
	p.Y += p.Vy
}
