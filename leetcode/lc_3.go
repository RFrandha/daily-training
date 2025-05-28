package leetcode

import (
	"log"
)

func L3() {
	// log.Println(lengthOfLongestSubstring("dvdf"))
	log.Println(lengthOfLongestSubstring("tmmzuxt"))
	// log.Println(lengthOfLongestSubstring("pwwkew"))
	// log.Println(lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	counter := 0
	maxCounter := 0
	for i, ss := range s {
		if v, ok := m[ss]; ok {
			if counter > maxCounter {
				maxCounter = counter
			}
			m[ss] = i
			counter = counter - v
		} else {
			m[ss] = i
			counter++
		}
	}
	if counter > maxCounter {
		return counter
	}
	return maxCounter
}
