package leetcode

import "log"

func L2364() {
	log.Println("Count Bad Pairs")
	log.Println(countBadPairs([]int{1, 2, 3, 4, 5}))
	// -1 -1 -1 -1 -1
	log.Println(countBadPairs([]int{4, 1, 3, 3}))
	// -4 0 -1 0
	// i - nums[i] = j - nums[j] good pair
}

func countBadPairs(nums []int) int64 {
	badCounter := 0
	goodDict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := i - nums[i]
		badCounter += i - goodDict[diff]
		goodDict[diff]++
	}

	return int64(badCounter)
}
