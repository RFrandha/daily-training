package leetcode

import "log"

func L1352() {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	log.Println(obj.GetProduct(2))
	log.Println(obj.GetProduct(3))
	log.Println(obj.GetProduct(5))
	obj.Add(8)
	log.Println(obj.GetProduct(2))
}

//sol 1
// type ProductOfNumbers struct {
//     nums    []int
// 	lastIdx int
// }

// func Constructor() ProductOfNumbers {
//     return ProductOfNumbers{
// 		nums:    make([]int, 40000),
// 		lastIdx: 0,
// 	}
// }

// func (this *ProductOfNumbers) Add(num int)  {
//     this.nums[this.lastIdx] = num
// 	this.lastIdx++
// }

// func (this *ProductOfNumbers) GetProduct(k int) int {
//     counter := 1
// 	for i := this.lastIdx - k; i < this.lastIdx; i++ {
// 		if this.nums[i] == 0 {
// 			return 0
// 		}
// 		counter *= this.nums[i]
// 	}
// 	return counter
// }

type ProductOfNumbers struct {
	nums []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		nums: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.nums = []int{1}
	} else {
		this.nums = append(this.nums, this.nums[len(this.nums)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.nums) {
		return 0
	}
	return this.nums[len(this.nums)-1] / this.nums[len(this.nums)-1-k]
}
