package leetcode

import "log"

func L2342() {
	log.Println(maximumSum([]int{18, 43, 36, 13, 7}))
}

func maximumSum(nums []int) int {
	maxSum := -1
	dictSum := make(map[int]int)
	for _, num := range nums {
		sumDigit := 0
		nowNum := num

		for nowNum > 0 {
			sumDigit += nowNum % 10
			nowNum = nowNum / 10
		}

		if dictSum[sumDigit] != 0 {
			maxSum = max(maxSum, dictSum[sumDigit]+num)
		}

		if num > dictSum[sumDigit] {
			dictSum[sumDigit] = num
		}
	}
	return maxSum
}
