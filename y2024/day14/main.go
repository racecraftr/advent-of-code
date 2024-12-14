package main

import (
	"adventOfCode/util"
	"adventOfCode/util/arrays"
	"adventOfCode/util/conv"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
)

//go:embed in.txt
var input string

type Robot struct {
	px, py int
	vx, vy int
}

func (r *Robot) update(seconds, maxX, maxY int) {
	ux := r.px + seconds*r.vx
	uy := r.py + seconds*r.vy
	if ux < 0 {
		n := (-ux / maxX) + 1
		ux += n * maxX
	}
	if uy < 0 {
		n := (-uy / maxY) + 1
		uy += n * maxY
	}
	ux %= maxX
	uy %= maxY
	r.px = ux
	r.py = uy
}

func (r *Robot) asPoint() point {
	return point{(*r).px, (*r).py}
}

func parseInput(input string) []*Robot {
	re := regexp.MustCompile("-*\\d+")
	return arrays.Map(util.Lines(input), func(line string) *Robot {
		parts := re.FindAllString(line, -1)
		px := conv.ToInt(parts[0])
		py := conv.ToInt(parts[1])
		vx := conv.ToInt(parts[2])
		vy := conv.ToInt(parts[3])

		return &Robot{px, py, vx, vy}
	})
}
func part1(input string) string {
	maxX, maxY := 101, 103
	midX, midY := maxX/2, maxY/2
	println(midX, midY)
	quadCount := make([]int, 4)
	robots := parseInput(input)
	for _, r := range robots {
		r.update(100, maxX, maxY)
		if r.px == midX || r.py == midY {
			fmt.Printf("%v is in middle\n", *r)
			continue
		}
		isTop := r.py < midY
		isLeft := r.px < midX
		if isTop && isLeft {
			fmt.Printf("%v goes in quad 1\n", *r)
			quadCount[0]++
		}
		if isTop && !isLeft {
			fmt.Printf("%v goes in quad 2\n", *r)
			quadCount[1]++
		}
		if !isTop && isLeft {
			fmt.Printf("%v goes in quad 3\n", *r)
			quadCount[2]++
		}
		if !isTop && !isLeft {
			fmt.Printf("%v goes in quad 4\n", *r)
			quadCount[3]++
		}
	}
	return conv.ToString(arrays.Product(quadCount))
}

type point [2]int

// render tree as test.
func render(points map[point]bool, maxX, maxY int) {
	mtx := make([][]rune, maxY)
	for i := range mtx {
		mtx[i] = make([]rune, maxX)
		for j := range mtx[i] {
			mtx[i][j] = ' '
		}
	}
	for p := range points {
		x, y := p[0], p[1]
		mtx[y][x] = '█'
	}
	for _, ln := range mtx {
		fmt.Println(string(ln))
	}
}

// saveImg saves the tree as an image. For funsies.
func saveImg(points map[point]bool, maxX, maxY int) {

	// 101	16	245
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	purp := color.RGBA{R: 101, G: 16, B: 245, A: 0xff}
	for y := range maxY {
		for x := range maxX {
			if points[point{x, y}] {
				img.Set(x, y, purp)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}

	f, _ := os.Create("day14.png")
	err := png.Encode(f, img)
	if err != nil {
		return
	}
}

func part2(input string) string {
	maxX, maxY := 101, 103
	robots := parseInput(input)
	seconds := 0
	for {
		points := map[point]bool{}
		for _, r := range robots {
			r.update(1, maxX, maxY)
		}
		seconds++
		for _, r := range robots {
			points[r.asPoint()] = true
		}
		if len(points) == len(robots) {
			//saveImg(points, maxX, maxY) // uncomment this line to save the tree as an image :)
			break
		}
	}
	return conv.ToString(seconds)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
