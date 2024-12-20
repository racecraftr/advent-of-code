package mathy

import (
	"testing"
)

func Test1(t *testing.T) {
	dist := ManhattanDist(0, 0, 1, 1)
	if dist != 2 {
		t.Errorf("Distance should be %v, was actually %v\n", 2, dist)
	}
}

func Test2(t *testing.T) {
	dist := ManhattanDist(-1, 1, 1, -1)
	if dist != 4 {
		t.Errorf("Distance should be %v, was actually %v\n", 4, dist)
	}
}

func TestIntAbs(t *testing.T) {
	n := IntAbs(-40 + 10)
	if n != 30 {
		t.Errorf("Expected %v, got %v", 30, n)
	}
}
