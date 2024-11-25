package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const INPUT = "reyedfim"

func part1() {
	res := ""
	i := 0
	for len(res) < 8 {
		hstr := hashMD5(INPUT + fmt.Sprintf("%d", i))
		if strings.HasPrefix(hstr, "00000") {
			res += string(hstr[5])
			println(hstr)
			println("found! res = " + res)
		}
		i ++
	}
	
	println(res)
}

func part2() {
	res := make([]byte, 8)
	for i := range res {
		res[i] = 0
	}
	count := 0
	i := 0
	for count < 8 {
		hstr := hashMD5(INPUT + fmt.Sprintf("%d", i))
		if strings.HasPrefix(hstr, "00000") {
			println(hstr)
			pos := int(hstr[5] - '0')
			ch := hstr[6]
			if pos >= 8 {
				goto loopend
			} else {
				if res[pos] == 0 {
					res[pos] = ch
					count ++
					println("res = " + string(res))
				}
			}
		}
		loopend:
		i ++
	}
	println(string(res))
}

func hashMD5(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}

func main() {
	part2()
}