package main

import "math"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left := -1
	m := map[int32]int{}

	for i, c := range s {
		if idx, ok := m[c]; ok {
			left = int(math.Max(float64(left), float64(idx)))
		}
		m[c] = i
		maxLen = int(math.Max(float64(maxLen), float64(i-left)))
	}
	return maxLen
}

func main() {
	ret := lengthOfLongestSubstring("pwwkew")
	println(ret)
}
