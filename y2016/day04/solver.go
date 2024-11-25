package main

import "sort"

type room struct {
	nameParts[]string
	sectorID int
	checksum string
}

type countedRune struct {
	r rune
	count int
}

func (r *room) isValid() bool {
	charCount := map[rune]int{}
	for _, part := range r.nameParts {
		for _, c := range part {
			charCount[c]++
		}
	}
	
	var counts []*countedRune

	for r, count := range charCount {
		counts = append(counts, &countedRune{r, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count != counts[j].count {
			return counts[i].count > counts[j].count
		}
		return counts[i].r < counts[j].r
	})

	checksum := ""
	for _, cr := range counts {
		checksum += string(cr.r)
	}

	return r.checksum == checksum
}