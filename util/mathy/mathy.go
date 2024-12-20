package mathy

func ManhattanDist(x1, y1, x2, y2 int) int {
	return IntAbs(x2-x1) + IntAbs(y2-y1)
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
