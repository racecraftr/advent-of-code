package grid

// Transpose from https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func Transpose[T any](grid [][]T) [][]T {
	xl := len(grid[0])
	yl := len(grid)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = grid[j][i]
		}
	}
	return result
}

// IsValidPos makes sure that a position in the grid is valid.
func IsValidPos[T any](grid [][]T, x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func Flatten[T any](grid [][]T) []T {
	flattened := make([]T, len(grid)*len(grid[0]))
	for i, arr := range grid {
		for j, v := range arr {
			flattened[i*len(arr)+j] = v
		}
	}
	return flattened
}

// Point stores itself as [x, y].
type Point [2]int

func (p Point) Add(other Point) Point {
	return Point{p[0] + other[0], p[1] + other[1]}
}

func At[T any](grid [][]T, p Point) (bool, T) {
	if IsValidPos(grid, p[0], p[1]) {
		return true, grid[p[0]][p[1]]
	}
	return false, grid[0][0]
}
