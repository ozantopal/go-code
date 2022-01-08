package main

import "fmt"

func main() {
	result := addBinary("1010", "1101001")
	fmt.Println(result)
}

func addBinary(a string, b string) string {
	arrA, arrB := []rune(reverse(a)), []rune(reverse(b))
	if len(arrA) < len(arrB) {
		arrA, arrB = arrB, arrA
	}

	carry := 0
	for i := 0; i < len(arrA); i++ {
		nextDigit := convert(int(arrA[i])) + carry
		if len(arrB) > i {
			nextDigit += convert(int(arrB[i]))
		}

		carry = nextDigit / 2
		nextDigit = nextDigit % 2

		if nextDigit == 0 {
			arrA[i] = 48
		} else if nextDigit == 1 {
			arrA[i] = 49
		}
	}
	if carry != 0 {
		arrA = append(arrA, rune(49))
	}

	return reverse(string(arrA))
}

func convert(a int) int {
	if a == 49 {
		return 1
	}
	return 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
