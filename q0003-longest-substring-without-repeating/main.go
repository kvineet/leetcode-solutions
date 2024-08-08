package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	var track map[rune]int = make(map[rune]int)
	var minPos, maxLen int = -1, 0

	for i, v := range s {
		val, ok := track[v]
		if ok && minPos < val {
			minPos = val
			track[v] = i
		}
		track[v] = i
		if maxLen < i-minPos {
			maxLen = i - minPos
		}
	}
	return maxLen
}
