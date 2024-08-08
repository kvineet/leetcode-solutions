package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type arg struct {
		num    []int
		target int
	}

	tests := []struct {
		arg  arg
		want []int
	}{
		{
			arg:  arg{num: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			arg:  arg{num: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			arg:  arg{num: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arg.num), func(t *testing.T) {
			got := twoSum(tt.arg.num, tt.arg.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
