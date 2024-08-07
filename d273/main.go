package main

import (
	"bytes"
	"slices"
	"strings"
)

var unitPlaceStr = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
var teenPlaceStr = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var decaPlaceStr = []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var hundredStr = "Hundred"
var groupStr = []string{"", " Thousand", " Million", " Billion"}

func main() {

}

// 1234567
// 1 234 567

func numberToWords(a int) string {
	if a == 0 {
		return "Zero"
	}
	var words bytes.Buffer
	groups := groupDigits(a)
	groupLen := len(groups)
	slices.Reverse(groups)
	for i, v := range groups {
		wordGroup := strings.TrimSpace(groupToWords(v))
		if wordGroup == "" {
			continue
		}
		words.WriteString(wordGroup)
		words.WriteString(groupStr[groupLen-i-1])
		words.WriteString(" ")
	}
	return strings.Trim(words.String(), " ")
}

func groupToWords(a int) string {
	if a > 999 {
		panic("group needs to be less than Thousand")
	}
	var result bytes.Buffer
	rem := a / 100
	a = a % 100
	if rem > 0 {
		result.WriteString(unitPlaceStr[rem-1])
		result.WriteString(" ")
		result.WriteString(hundredStr)
		result.WriteString(" ")
	}
	if a < 21 && a > 9 {
		result.WriteString(teenPlaceStr[a-10])
	} else {
		rem := a / 10
		a = a % 10
		if rem > 0 {
			result.WriteString(decaPlaceStr[rem-2])
			result.WriteString(" ")
		}
		if a > 0 {
			result.WriteString(unitPlaceStr[a-1])
		}
	}
	return result.String()
}

func groupDigits(a int) []int {
	result := []int{}
	for a != 0 {
		rem := a % 1000
		result = append(result, rem)
		a = a / 1000
	}
	return result
}
