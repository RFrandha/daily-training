package leetcode

func LC55(nums []int) bool {
	//lastIdx := len(nums) - 1
	return canJump(nums)
}

func checkLastIdx(currNum int, nums []int, currIdx, lastIdx int) bool {
	if currIdx+currNum >= lastIdx {
		return true
	}
	for i := 1; i <= currNum; i++ {
		if checkLastIdx(nums[currIdx+i], nums, currIdx+i, lastIdx) {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	limit := 0
	for i := 0; i < len(nums); i++ {
		if i > limit {
			return false
		}
		if i+nums[i] > len(nums)-1 {
			return true
		}
		if i+nums[i] > limit {
			limit = i + nums[i]
		}
	}
	return true
}
