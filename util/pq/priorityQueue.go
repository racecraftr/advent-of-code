// Package pq is responsible for storing a Priority Queue backed by a Min Heap.
// It does not use the standard heap package in go, rather instead creating a normal heap.
//
// A priority queue here can store any value.
package pq

type PriorityQueue[T any] struct {
	heap []T
	less func(T, T) bool
}

// NewPQ should take a function less. The "less" function acts as a way for sifting,
// in which it checks if the left parameter compared to the right parameter.
func NewPQ[T any](less func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: make([]T, 0),
		less: less,
	}
}

// IntPQ is just a basic PriorityQueue for ints.
func IntPQ() *PriorityQueue[int] {
	return NewPQ(func(a int, b int) bool {
		return a < b
	})
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (left, right int) {
	return 2*i + 1, 2*i + 2
}

func (pq *PriorityQueue[T]) Push(t T) {
	pq.heap = append(pq.heap, t)
	idx := len(pq.heap) - 1
	for idx > 0 && pq.less(pq.heap[idx], pq.heap[parent(idx)]) {
		p := parent(idx)
		pq.heap[idx], pq.heap[p] = pq.heap[p], pq.heap[idx]
		idx = p
	}
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue[T]) Pop() T {
	if pq.Len() <= 0 {
		panic("cannot pop from empty heap")
	}
	n := len(pq.heap)
	pq.heap[0], pq.heap[n-1] = pq.heap[n-1], pq.heap[0]
	popped := pq.heap[n-1]
	pq.heap = pq.heap[:n-1]

	idx := 0
	for idx < len(pq.heap) {
		left, right := children(idx)

		if left >= len(pq.heap) && right >= len(pq.heap) {
			break
		}

		minIdx := idx

		if left < len(pq.heap) && pq.less(pq.heap[left], pq.heap[minIdx]) {
			minIdx = left
		}
		if right < len(pq.heap) && pq.less(pq.heap[right], pq.heap[minIdx]) {
			minIdx = right
		}

		if minIdx == idx {
			break
		}

		pq.heap[minIdx], pq.heap[idx] = pq.heap[idx], pq.heap[minIdx]
		idx = minIdx
	}

	return popped
}
