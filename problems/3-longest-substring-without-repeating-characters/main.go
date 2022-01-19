package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)

	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	right, left, result := 0, 0, 0
	indexes := make(map[byte]int, len(s))

	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		result = max(result, left-right)
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
