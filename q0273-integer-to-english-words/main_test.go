package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupToWords(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 5, want: "Five",
		},
		{
			input: 51, want: "Fifty One",
		},
		{
			input: 12, want: "Twelve",
		},
		{
			input: 154, want: "One Hundred Fifty Four",
		},
		{
			input: 414, want: "Four Hundred Fourteen",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := groupToWords(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}

}
func TestGroupDigits(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{
			input: 5, want: []int{5},
		},
		{
			input: 51, want: []int{51},
		},
		{
			input: 154, want: []int{154},
		},
		{
			input: 2535, want: []int{535, 2},
		},
		{
			input: 1234343453, want: []int{453, 343, 234, 1},
		},
		{
			input: 50868, want: []int{868, 50},
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := groupDigits(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 5, want: "Five",
		},
		{
			input: 51, want: "Fifty One",
		},
		{
			input: 154, want: "One Hundred Fifty Four",
		},
		{
			input: 2535, want: "Two Thousand Five Hundred Thirty Five",
		},
		{
			input: 1234343453, want: "One Billion Two Hundred Thirty Four Million Three Hundred Forty Three Thousand Four Hundred Fifty Three",
		},
		{
			input: 1000000, want: "One Million",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := numberToWords(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
