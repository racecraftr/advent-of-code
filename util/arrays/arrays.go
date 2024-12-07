package arrays

func ValidPosition[T any](arr []T, i int) bool {
	return i >= 0 && i < len(arr)
}

func Map[T, U any](arr []T, mapFunc func(T) U) []U {
	res := make([]U, len(arr))
	for i, t := range arr {
		res[i] = mapFunc(t)
	}
	return res
}

func Filter[T any](arr []T, filterFunc func(T) bool) []T {
	res := make([]T, len(arr))
	i := 0
	for _, t := range arr {
		if filterFunc(t) {
			res[i] = t
			i++
		}
	}
	return res[:i]
}

func All[T any](arr []T, f func(T) bool) bool {
	for _, t := range arr {
		if !f(t) {
			return false
		}
	}
	return true
}

func Any[T any](arr []T, f func(T) bool) bool {
	for _, t := range arr {
		if f(t) {
			return true
		}
	}
	return false
}

func Sum[T int | uint | float64](arr []T) T {
	var sum T = 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func Product[T int | uint | float64](arr []T) T {
	var product T = 1
	for _, n := range arr {
		if n == 0 {
			return 0
		}
		product *= n
	}
	return product
}
