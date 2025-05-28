package leetcode

import "log"

func LC4() {
	log.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	//log.Println(findMedianSortedArrays([]int{}, []int{1}))
}

//
//1 3 4
//2 5 6
//
//3
//2
//1
//
//3 4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	medIdx1 := ((len(nums1) + len(nums1)) / 2) + 1
	isEven := false
	if (len(nums1)+len(nums2))%2 == 0 {
		isEven = true
	}
	counter := 1

	nums1Pointer := 0
	nums2Pointer := 0

	for i := 0; i < len(nums1)+len(nums2); i++ {
		if counter == medIdx1 {
			firstNum := 0.0
			if len(nums2)-1 < nums2Pointer {
				firstNum = float64(nums1[nums1Pointer])
				nums1Pointer++
			} else if len(nums1)-1 < nums1Pointer {
				firstNum = float64(nums2[nums2Pointer])
				nums2Pointer++
			}
			if firstNum == 0.0 && nums1[nums1Pointer] < nums2[nums2Pointer] {
				firstNum = float64(nums1[nums1Pointer])
				nums1Pointer++
			} else if firstNum == 0.0 {
				firstNum = float64(nums2[nums2Pointer])
				nums2Pointer++
			}
			if isEven {
				secondNum := 0.0
				if len(nums2)-1 < nums2Pointer {
					secondNum = float64(nums1[nums1Pointer])
				} else if len(nums1)-1 < nums1Pointer {
					secondNum = float64(nums2[nums2Pointer])
				}
				if secondNum == 0.0 && nums1[nums1Pointer] < nums2[nums2Pointer] {
					secondNum = float64(nums1[nums1Pointer])
				} else if secondNum == 0.0 {
					secondNum = float64(nums2[nums2Pointer])
				}
				return (firstNum + secondNum) / 2.0
			} else {
				return firstNum
			}
		}
		if len(nums1) > 0 && nums1[nums1Pointer] < nums2[nums2Pointer] {
			nums1Pointer++
		} else {
			nums2Pointer++
		}
		counter++
	}
	return 0
}
