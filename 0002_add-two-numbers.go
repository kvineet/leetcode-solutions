package main

// ListNode defines a single linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry < 1 {
			return nil
		}
		return &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	var l1d, l2d *ListNode
	result := carry
	if l1 != nil {
		result = result + l1.Val
		l1d = l1.Next
	}
	if l2 != nil {
		result = result + l2.Val
		l2d = l2.Next
	}
	newCarry := result / 10
	return &ListNode{
		Val:  result % 10,
		Next: addTwoNumbersHelper(l1d, l2d, newCarry),
	}
}
