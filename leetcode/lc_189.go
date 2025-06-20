package leetcode

import "log"

func LC189() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	log.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	r := k % n

	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	leftPtr := 0
	rightPtr := len(nums) - 1

	for leftPtr < rightPtr {
		swap(leftPtr, rightPtr)
		leftPtr++
		rightPtr--
	}

	leftPtr = 0
	rightPtr = r - 1
	for leftPtr < rightPtr {
		swap(leftPtr, rightPtr)
		leftPtr++
		rightPtr--
	}

	leftPtr = r
	rightPtr = n - 1
	for leftPtr < rightPtr {
		swap(leftPtr, rightPtr)
		leftPtr++
		rightPtr--
	}
}
