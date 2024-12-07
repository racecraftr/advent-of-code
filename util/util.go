package util

import (
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var WhiteSpaceRegex = regexp.MustCompile("\\s+")

// Transpose from https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func Transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func SplitSpace(input string) []string {
	return WhiteSpaceRegex.Split(input, -1)
}
