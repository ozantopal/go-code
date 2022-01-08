package main

import "fmt"

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	lenS, isWordBegin := len(s), false
	result := 0

	for i := lenS - 1; i >= 0; i-- {
		if isWordBegin && s[i] == 32 {
			return result
		}
		if !isWordBegin && s[i] == 32 {
			continue
		}

		result++
		isWordBegin = true
	}

	return result
}
