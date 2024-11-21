package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	path := fmt.Sprintf("y%d/day%02d/in.txt", year, day)
	absPath, err := filepath.Abs(path)
	check(err)

	file, err := os.Open(absPath)
	check(err)
	defer check(file.Close())

	lines, i := make([]string, 2048), 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines[i] = line
		fmt.Println(line)
		i++
	}

	return lines[:i]
}
