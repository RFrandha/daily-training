package leetcode

import (
	"log"
	"time"
)

func L2467() {
	now := time.Now()
	eod := now.Truncate(24 * time.Hour).Add(-time.Second)
	ttl := eod.Sub(now)
	log.Println(ttl.String())
	log.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	return 0
}
