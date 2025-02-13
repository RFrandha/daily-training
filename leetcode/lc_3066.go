package leetcode

import (
	"container/heap"
	"log"
	"slices"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func L3066() {
	log.Println(minOperations([]int{1, 1, 2, 4, 9}, 20))
	log.Println(minOperations2([]int{1, 1, 2, 4, 9}, 20))
}

func minOperations(nums []int, k int) int {
	counter := 0
	nums = slices.Sorted(slices.Values(nums))
	for nums[0] < k {
		val := nums[0]*2 + nums[1]
		if len(nums) == 2 {
			if val < k {
				break
			} else {
				counter++
				return counter
			}
		}

		nums = nums[2:]
		nums = slices.Insert(nums, sort.SearchInts(nums, val), val)
		counter++
	}
	return counter
}

func minOperations2(nums []int, k int) int {
	counter := 0
	var h IntHeap = nums

	heap.Init(&h)

	for h.Len() > 1 {
		if h[0] >= k {
			break
		}

		val := heap.Pop(&h).(int)*2 + heap.Pop(&h).(int)
		heap.Push(&h, val)
		counter++
	}
	return counter
}
