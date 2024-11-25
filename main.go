package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	const N = 1000
	matA := initMatrix(N, 1)
	matB := initMatrix(N, 2)
	res := initMatrix(N, 0)
	
	start := currentTime()
	for i := 0; i < N; i ++ {
		for j := 0; j < N; j ++ {
			for k := 0; k < N; k ++ {
				res[i][j] += matA[i][k] * matB[k][j]
			}
		}
	}
	end := currentTime()
	
	fmt.Printf("Time elapsed: %v seconds\n", float64(end - start) / float64(time.Second))
}

func currentTime() int64{
	return time.Now().UnixNano()
}

func initMatrix(size, init int) [][]int{
	mat := make([][]int, size)
	for i := range size {
		mat[i] = make([]int, size)
		for j := range size {
			mat[i][j] = init
		}
	}
	return mat
}

func isPrime(n int) bool {
	if (n < 2) {
		return false
	}
	lim := int(math.Sqrt(float64(n)))
	for i := 2; i <= lim; i ++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
