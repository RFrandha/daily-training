package leetcode

import (
	"log"
	"slices"
)

func L3066() {
	log.Println(minOperations([]int{999999999, 999999999, 999999999}, 1000000000))
}

func minOperations(nums []int, k int) int {
	counter := 0
	nums = slices.Sorted(slices.Values(nums))
	for nums[0] < k {
		counter++
		val := nums[0]*2 + nums[1]
		nums = nums[2:]
		for i, v := range nums {

		}
	}
	return counter
}
