package leetcode

import (
	"container/heap"
	"log"
	"sort"
)

func LC3362() {
	//log.Println(maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}))
	//log.Println(maxRemoval([]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}))
	//log.Println(maxRemoval([]int{1, 2, 3, 4}, [][]int{{0, 3}}))
	//log.Println(maxRemoval([]int{0, 3}, [][]int{{0, 1}, {0, 0}, {0, 1}, {0, 1}, {0, 0}}))
	//log.Println(maxRemoval([]int{1, 2}, [][]int{{1, 1}, {0, 0}, {1, 1}, {1, 1}, {0, 1}, {0, 0}}))
	log.Println(maxRemoval([]int{0, 0, 1, 1, 0, 0}, [][]int{{2, 3}, {0, 2}, {3, 5}}))
}

//func maxRemoval(nums []int, queries [][]int) int {
//	sort.Slice(queries, func(i, j int) bool {
//		return queries[i][1]-queries[i][0] > queries[j][1]-queries[j][0] || queries[i][0] < queries[j][0] || queries[i][1] > queries[j][1]
//	})
//	log.Println(queries)
//
//	skipQuery := 0
//	for k := 0; k < len(queries); k++ {
//		changingQuery := false
//		for l := queries[k][0]; l <= queries[k][1]; l++ {
//			if nums[l] > 0 {
//				nums[l]--
//				changingQuery = true
//			}
//		}
//		if !changingQuery {
//			skipQuery++
//		}
//
//		zeroArray := true
//		for m := 0; m < len(nums); m++ {
//			if nums[m] > 0 {
//				zeroArray = false
//				break
//			}
//		}
//		//log.Println(nums)
//		if zeroArray {
//			return len(queries) - k - 1 + skipQuery
//		}
//	}
//	return -1
//}

func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})
	pq := &Heap{}
	heap.Init(pq)
	deltaArray := make([]int, len(nums)+1)
	operations := 0

	for i, j := 0, 0; i < len(nums); i++ {
		operations += deltaArray[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(pq, queries[j][1])
			j++
		}
		for operations < nums[i] && pq.Len() > 0 && (*pq)[0] >= i {
			operations += 1
			deltaArray[heap.Pop(pq).(int)+1] -= 1
		}
		if operations < nums[i] {
			return -1
		}
	}
	return pq.Len()
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
