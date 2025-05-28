package leetcode

import "log"

func LC40() {
	//log.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	log.Println(combinationSum2([]int{1, 1, 2, 5, 6, 7, 10}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	for i, val := range candidates {
		if val > target {
			continue
		} else if val == target {
			res = append(res, []int{val})
		} else {
			getCombination(candidates[i+1:], val, target, []int{val}, &res)
		}
	}

	return res
}

func getCombination(candidates []int, currVal, target int, currArr []int, res *[][]int) {
	for j, val := range candidates {
		if currVal+val > target {
			continue
		} else if currVal+val == target {
			*res = append(*res, append(currArr, val))
		} else {
			getCombination(candidates[j+1:], currVal+val, target, append(currArr, val), res)
		}
	}
}
