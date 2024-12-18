package pq

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPq(t *testing.T) {
	pq := NewPQ(func(a, b int) bool {
		return a < b
	})
	unsorted := []int{5, 6, -1, 3}
	sorted := []int{-1, 3, 5, 6}
	for _, v := range unsorted {
		pq.Push(v)
	}
	unsorted = []int{}
	for pq.Len() > 0 {
		unsorted = append(unsorted, pq.Pop())
	}

	if len(unsorted) != len(sorted) {
		t.Errorf("Unsorted has length %v while sorted has length %v", len(unsorted), len(sorted))
	}
	for i, v := range unsorted {
		if v != sorted[i] {
			t.Errorf("Contents of unsorted and sorted are not the same")
		}
	}

	fmt.Printf("Unsorted = %v\n", unsorted)
}

func TestPq2(t *testing.T) {
	const l = 100
	pq := IntPQ()
	for range l {
		pq.Push(rand.Intn(100))
	}

	var sorted = make([]int, 0, l)
	for pq.Len() > 0 {
		sorted = append(sorted, pq.Pop())
	}

	if len(sorted) != l {
		t.Errorf("Sorted is not correct length")
	}

	for i := range len(sorted) - 1 {
		if sorted[i] > sorted[i+1] {
			t.Errorf("sorted[%v] = %v, sorted[%v] = %v", i, sorted[i], i+1, sorted[i+1])
		}
	}

	fmt.Printf("Sorted = %v\n", sorted)
}
