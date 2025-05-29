package leetcode

import (
	"log"
)

func LC41() {
	log.Println(firstMissingPositive([]int{1, 2, 0}))
	log.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	log.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	// log.Println(firstMissingPositive([]int{1, 2, 3}))
	// log.Println(firstMissingPositive([]int{1}))
	// log.Println(firstMissingPositive([]int{}))
	// log.Println(firstMissingPositive([]int{2}))
}
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
