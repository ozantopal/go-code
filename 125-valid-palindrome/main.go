package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lenS := len(s)
	clearedChars := make([]byte, lenS)
	index := 0
	for i := 0; i < lenS; i++ {
		if isAcceptableChar(s[i]) {
			clearedChars[index] = s[i]
			index++
		}
	}

	clearedStr := string(clearedChars[0:index])
	return clearedStr == reverse(clearedStr)
}

func isAcceptableChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
