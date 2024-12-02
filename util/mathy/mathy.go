package mathy

import "math"

func ManhattanDist(x1, y1, x2, y2 int) int {
	return IntAbs((x2 - x1) + (y2 - y1))
}

func IntAbs(n int) int {
	return int(math.Abs(float64(n)))
}
