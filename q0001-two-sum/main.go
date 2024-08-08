package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
	twoSum([]int{2, 7, 11, 15}, 9)
	twoSum([]int{2, 7, 11, 15}, 9)
}

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
