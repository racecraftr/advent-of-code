package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetContent(year, day int) string {
	path := fmt.Sprintf("y%d/day%02d/in.txt", year, day)
	absPath, err := filepath.Abs(path)
	check(err)

	bytes, err := os.ReadFile(absPath)
	check(err)
	content := string(bytes)
	return content

}

func GetLines(year, day int) []string {
	return strings.Split(GetContent(year, day), "\n")
}
