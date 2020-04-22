package problems

import (
	"math"
)

func minWindow(s string, t string) string {
	src := []rune(s)
	tSrc := []rune(t)
	r, l := 0, 0
	minLen := math.MaxInt32
	needs, window := make(map[rune]int), make(map[rune]int)
	match := 0
	start := -1
	for _, v := range tSrc {
		if _, ok := needs[v]; ok {
			needs[v] += 1
		} else {
			needs[v] = 1
		}
		window[v] = 0
	}

	for r < len(src) {
		// needs this rune
		c1 := src[r]
		if _, ok := needs[c1]; ok {
			window[c1]++
			if window[c1] == needs[c1] {
				//fmt.Println("match")
				match++
			}
		}
		r++
		//fmt.Println(needs, window)
		for match == len(needs) {
			if minLen > (r - l) {
				minLen = r - l
				start = l
				//fmt.Println("update")
			}

			c2 := src[l]
			if _, ok := needs[c2]; ok {
				window[c2] -= 1
				if window[c2] < needs[c2] {
					match--
				}
			}
			l++
		}

	}
	if start == -1 {
		return ""
	}
	return s[start : start+minLen]
}
