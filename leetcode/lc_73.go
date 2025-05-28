package leetcode

import "log"

func LC73() {
	log.Println(setZeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))
}

//00 01 02
//10 11 12
//20 21 22

func setZeroes(matrix [][]int) [][]int {
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = true
				colSet[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rowSet[i] {
				matrix[i][j] = 0
			}
			if colSet[j] {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
