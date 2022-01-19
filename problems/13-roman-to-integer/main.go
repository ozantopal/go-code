package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	numbs := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i, ch := range s {
		if i+1 < len(s) {
			if ch == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
				result -= numbs[ch]
				continue
			}
			if ch == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
				result -= numbs[ch]
				continue
			}
			if ch == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
				result -= numbs[ch]
				continue
			}
		}
		result += numbs[ch]
	}

	return result
}
