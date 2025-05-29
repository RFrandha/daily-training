package leetcode

import (
	"log"
	"slices"
)

func LC40() {
	log.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// log.Println(combinationSum2([]int{1, 1, 2, 5, 6, 7, 10}, 8))
	log.Println(combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 30))

}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	slices.Sort(candidates)
	getCombination(candidates, 0, target, []int{}, &res)
	return res
}

func getCombination(candidates []int, startIdx, target int, currArr []int, res *[][]int) {
	if target == 0 {
		comb := make([]int, len(currArr))
		copy(comb, currArr)
		*res = append(*res, comb)
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		if i > startIdx && candidates[i] == candidates[i-1] {
			continue
		}

		if candidates[i] > target {
			break
		}

		currArr = append(currArr, candidates[i])
		getCombination(candidates, i+1, target-candidates[i], currArr, res)
		currArr = currArr[:len(currArr)-1]
	}
}
