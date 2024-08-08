package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SliceToListNode(l []int) *ListNode {
	var head, ptr *ListNode

	for _, v := range l {
		newNode := ListNode{Val: v, Next: nil}
		if ptr == nil {
			ptr = &newNode
			head = ptr
		} else {
			ptr.Next = &newNode
			ptr = ptr.Next
		}
	}
	return head
}
func ListNodeToSlice(l *ListNode) []int {
	var result []int
	if l == nil {
		return result
	}
	for {
		result = append(result, l.Val)
		l = l.Next
		if l == nil {
			break
		}
	}
	return result
}

func TestRoundTrip(t *testing.T) {
	l := []int{7, 0, 8}
	lst := SliceToListNode(l)
	assert.Equal(t, l, ListNodeToSlice(lst))
}

func TestAddTwoNumbers(t *testing.T) {
	type arg struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		arg  arg
		want []int
	}{
		{
			arg:  arg{l1: []int{2, 4, 3}, l2: []int{5, 6, 4}},
			want: []int{7, 0, 8},
		},
		{
			arg:  arg{l1: []int{0}, l2: []int{0}},
			want: []int{0},
		},
		{
			arg:  arg{l1: []int{9, 9, 9, 9, 9, 9, 9}, l2: []int{9, 9, 9, 9}},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for i, tt := range tests {
		t.Run("testcase_"+strconv.Itoa(i), func(t *testing.T) {
			got := addTwoNumbers(SliceToListNode(tt.arg.l1), SliceToListNode(tt.arg.l2))
			assert.Equal(t, tt.want, ListNodeToSlice(got))
		})
	}
}
