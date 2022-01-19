package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS < 2 {
		return s
	}

	begin, maxLen := 0, 1

	for i := 0; i < lenS; {
		if len(s)-i <= maxLen/2 {
			break
		}

		b, e := i, i
		for e < lenS-1 && s[e+1] == s[e] {
			e++
		}

		i = e + 1

		for e < lenS-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}

		newLen := e + 1 - b
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}
