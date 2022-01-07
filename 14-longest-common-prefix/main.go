package main

import "fmt"

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
	}
	i := 0
	var result []byte

	for {
		var nextCh byte = 0
		for _, str := range strs {
			if i >= len(str) {
				return string(result)
			}
			if nextCh == 0 {
				nextCh = str[i]
			} else {
				if nextCh != str[i] {
					return string(result)
				}
			}
		}

		i++
		result = append(result, nextCh)
	}
}
