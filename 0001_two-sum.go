package main

func twoSum(nums []int, target int) []int {
	var hashMap map[int]int = map[int]int{}
	var left int
	for i, v := range nums {
		left = target - v
		if val, ok := hashMap[left]; ok {
			return []int{val, i}
		}
		hashMap[v] = i
	}
	return []int{-1, -1}
}
