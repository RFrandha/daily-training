package leetcode

import (
	"log"
	"strings"
)

func L1910() {
	log.Println(removeOccurrences("axxxxyyyyb", "xy"))
}

//func removeOccurrences(s string, part string) string {
//	if strings.Contains(s, part) {
//		return removeOccurrences(strings.Replace(s, part, "", 1), part)
//	}
//	return s
//}

func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		s = strings.Replace(s, part, "", 1)
	}
	return s
}
