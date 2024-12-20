package maps

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Keys[K comparable, V any](mp map[K]V) []K {
	arr := make([]K, len(mp))
	i := 0
	for k := range mp {
		arr[i] = k
		i++
	}
	return arr
}

func Values[K comparable, V any](mp map[K]V) []V {
	arr := make([]V, len(mp))
	i := 0
	for _, v := range mp {
		arr[i] = v
		i++
	}
	return arr
}

func Entries[K comparable, V any](mp map[K]V) []*Entry[K, V] {
	arr := make([]*Entry[K, V], len(mp))
	i := 0
	for k, v := range mp {
		arr[i] = &Entry[K, V]{k, v}
		i++
	}
	return arr
}

func Copy[K comparable, V any](mp map[K]V) map[K]V {
	res := make(map[K]V)
	for k, v := range mp {
		res[k] = v
	}
	return res
}
