package leetcode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func L2() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	log.Println(addTwoNumbers(l1, l2))
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	result := &ListNode{0, nil}
// 	currentL1 := l1
// 	currentL2 := l2
// 	tempVal := 0
// 	ph := result

// 	for currentL1 != nil || currentL2 != nil || tempVal > 0 {

// 		if currentL1 != nil {
// 			tempVal += currentL1.Val
// 		}
// 		if currentL2 != nil {
// 			tempVal += currentL2.Val
// 		}

// 		ph.Next = &ListNode{tempVal % 10, nil}
// 		ph = ph.Next
// 		tempVal /= 10

// 		if currentL1 != nil {
// 			currentL1 = currentL1.Next
// 		}
// 		if currentL2 != nil {
// 			currentL2 = currentL2.Next
// 		}
// 	}

// 	return result.Next
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	currentL1 := l1
	currentL2 := l2
	tempVal := 0
	ph := result

	for currentL1 != nil || currentL2 != nil || tempVal > 0 {

		if currentL1 != nil {
			tempVal += currentL1.Val
			currentL1 = currentL1.Next
		}
		if currentL2 != nil {
			tempVal += currentL2.Val
			currentL2 = currentL2.Next
		}

		ph.Next = &ListNode{tempVal % 10, nil}
		ph = ph.Next
		tempVal /= 10
	}

	return result.Next
}
