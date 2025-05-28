package leetcode

import (
	"log"
)

func L9() {
	log.Println(isPalindrome(121))
	log.Println(isPalindrome(-121))
	log.Println(isPalindrome(10))
	log.Println(isPalindrome(-101))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverseX := 0
	originalX := x
	for x > 0 {
		reverseX = reverseX*10 + x%10
		x = x / 10
	}
	return originalX == reverseX
}
