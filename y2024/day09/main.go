package main

import (
    "adventOfCode/util"
    "adventOfCode/util/conv"
    _ "embed"
    "flag"
    "fmt"
)

//go:embed in.txt
var input string

func parseInput(input string) []int {
    current := 0
    res := make([]int, 0)
    for i, v := range input {
        n := int(v - '0')
        if i%2 == 0 {
            for range n {
                res = append(res, current)
            }
            current++
        } else {
            for range n {
                res = append(res, -1)
            }
        }
    }
    return res
}

func part1(input string) string {
    blocks := parseInput(input)
    i, j := 0, len(blocks)-1
    for i < j {
        for ; blocks[i] != -1; i++ {
        }
        for ; blocks[j] == -1; j-- {
        }
        blocks[i], blocks[j] = blocks[j], blocks[i]
    }
    sum := 0
    for idx, v := range blocks[:j+10] {
        if v != -1 {
            sum += idx * v
        }
    }
    return conv.ToString(sum)
}

func part2(input string) string {
    blocks := parseInput(input)

    findFile := func(pos int) (int, int) {
        if pos < 0 || pos >= len(blocks) || blocks[pos] == -1 {
            return -1, 0
        }

        id := blocks[pos]
        start := pos
        for ; start >= 0 && blocks[start] == id; start-- {
        }
        start++

        end := pos
        for ; end < len(blocks) && blocks[end] == id; end++ {
        }
        return start, end - start
    }

    findFreeSize := func(pos int) int {
        size := 0
        for i := pos; i < len(blocks) && blocks[i] == -1; i++ {
            size++
        }
        return size
    }

    maxId := -1
    for _, v := range blocks {
        maxId = max(v, maxId)
    }

    for id := maxId; id >= 0; id-- {
        filePos, fs := -1, -1
        for i := len(blocks) - 1; i >= 0; i-- {
            if blocks[i] == id {
                filePos, fs = findFile(i)
                break
            }
        }

        if filePos == -1 {
            continue
        }

        putPos := -1
        for i := range filePos {
            if blocks[i] == -1 {
                emptySize := findFreeSize(i)
                if emptySize >= fs {
                    putPos = i
                    break
                }
                i += emptySize - 1
            }
        }

        if putPos != -1 && putPos < filePos {
            for i := range fs {
                blocks[putPos+i], blocks[filePos+i] =
                    blocks[filePos+i], blocks[putPos+i]
            }
        }
    }

    sum := 0
    for i, v := range blocks {
        if v != -1 {
            sum += i * v
        }
    }
    return conv.ToString(sum)
}

// some testing stuff :P
func test1() {
    testStr := "2333133121414131402"
    fmt.Println("Test:", part1(testStr))
}

func test2() {
    testStr := "2333133121414131402"
    fmt.Println("Test:", part2(testStr))
}

func main() {
    var part int
    flag.IntVar(&part, "part", 1, "part 1 or 2")
    flag.Parse()
    fmt.Println("Running part", part)
    if part == 1 {
        test1()
        ans := part1(input)
        util.CopyToClipboard(fmt.Sprintf("%v", ans))
        fmt.Println("Output:", ans)
    } else {
        test2()
        ans := part2(input)
        util.CopyToClipboard(fmt.Sprintf("%v", ans))
        fmt.Println("Output:", ans)
    }
}
