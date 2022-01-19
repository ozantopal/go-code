package main

import "fmt"

func main() {
	haystack, needle := "", "a"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if haystack == needle {
		return 0
	}
	if len(needle) > len(haystack) || len(haystack) == 0 {
		return -1
	}

	res, pos := "", 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if i+len(needle) <= len(haystack) {
				res = haystack[i : i+(len(needle))]
				if res == needle {
					pos = i
					return pos
				}
			} else {
				return -1
			}
		} else {
			continue
		}
	}
	return -1
}
