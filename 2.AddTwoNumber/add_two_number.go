package main

import "fmt"

// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func main() {
	fmt.Println(addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. Brute Force
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = &ListNode{}
	var temp *ListNode = res
	var carry int = 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		temp.Next = &ListNode{Val: sum % 10}
		temp = temp.Next
	}
	return res.Next
}
