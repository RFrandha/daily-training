package leetcode

import (
	"log"
	"strconv"
	"strings"
)

func L3174() {
	log.Println("Clear Digits")
	log.Println(clearDigits("3acb34"))
	log.Println(len([]int{1, 2, 3, 4, 5}))

	a := []int{1, 2, 3, 4, 5}
	log.Println(len(a))
	log.Println(a[1:2])
}

func clearDigits(s string) string {
	res := ""
	for _, v := range strings.Split(s, "") {
		if _, err := strconv.ParseInt(v, 0, 10); err != nil {
			res += v
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return res
}
