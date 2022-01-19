package main

import (
	"fmt"
)

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	counter := len(digits)
	carry := 0
	result := make([]int, counter+1)

	for counter > 0 {
		nextDigit := digits[counter-1] + carry
		if counter == len(digits) {
			nextDigit++
		}
		carry = nextDigit / 10
		result[counter] = nextDigit % 10
		counter--
	}

	if carry > 0 {
		result[0] = carry
		return result
	} else {
		return result[1:]
	}
}
