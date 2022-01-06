package main

import "fmt"

func main() {
	x := 1222
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	countOfDigit := countDigits(x)
	isOdd := countOfDigit%2 == 1
	var digits []int

	for i := 0; i < countOfDigit; i++ {
		if countOfDigit/2 > i {
			digits = append(digits, int(getNextDigit(x)))
			x = x / 10
		} else if isOdd && countOfDigit/2 == i {
			x = x / 10
		} else {
			if rune(digits[len(digits)-1]) != getNextDigit(x) {
				return false
			}
			digits = digits[:len(digits)-1]
			x = x / 10
		}
	}

	return true
}

func countDigits(numb int) int {
	if numb/10 == 0 {
		return 1
	}
	return 1 + countDigits(numb/10)
}

func getNextDigit(x int) rune {
	if x == 0 {
		return -1
	}

	return rune(x % 10)
}
