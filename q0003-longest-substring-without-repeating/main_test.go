package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type arg struct {
		s string
	}
	tests := []struct {
		arg  arg
		want int
	}{
		{
			arg:  arg{s: "abcabcbb"},
			want: 3,
		},
		{
			arg:  arg{s: "bbbbbbb"},
			want: 1,
		},
		{
			arg:  arg{s: "pwwkew"},
			want: 3,
		},
		{
			arg:  arg{s: "a"},
			want: 1,
		},
		{
			arg:  arg{s: "au"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run("testcase_"+tt.arg.s, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.arg.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
