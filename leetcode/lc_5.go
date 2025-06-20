package leetcode

import (
	"log"
	"strings"
)

func LC5() {
	//log.Println(checkPalindrome("a"))
	//log.Println(longestPalindrome("a"))
	log.Println(longestPalindrome("abbc"))
}

func longestPalindrome(s string) string {
	splitString := strings.Split(s, "")
	size := len(splitString)

	for size > 0 {
		for i := 0; i < len(s)-size; i++ {
			if checkPalindrome(strings.Join(splitString[i:i+size+1], "")) {
				return strings.Join(splitString[i:i+size+1], "")
			}
		}
		size--
	}
	if len(s) > 0 {
		return splitString[0]
	} else {
		return ""
	}

	//for i := 0; i < len(s); i++ {
	//	for j := len(s) - 1; j >= i; j-- {
	//		if checkPalindrome(strings.Join(splitString[i:j+1], "")) {
	//			size := j + 1 - i
	//			if longest < size {
	//				longest = size
	//			}
	//			lenStrMap[size] = strings.Join(splitString[i:j+1], "")
	//		}
	//	}
	//}
	//return lenStrMap[longest]
}

//func longestPalindrome(s string) string {
//	splitString := strings.Split(s, "")
//	for i := 0; i < len(s); i++ {
//		for j := len(s) - 1; j >= i; j-- {
//			if checkPalindrome(strings.Join(splitString[i:j+1], "")) {
//				return strings.Join(splitString[i:j+1], "")
//			}
//		}
//	}
//	return ""
//}

func checkPalindrome(s string) bool {
	leftPoint := 0
	rightPoint := len(s) - 1

	for leftPoint < rightPoint {
		if s[leftPoint] == s[rightPoint] {
			leftPoint++
			rightPoint--
			continue
		} else {
			return false
		}
	}
	return true
}
