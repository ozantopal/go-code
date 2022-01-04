package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var isNegative bool = false

	if x < 0 {
		x *= -1
		isNegative = true
	}
	s := strconv.Itoa(x)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if isNegative {
	}
	res, err := strconv.Atoi(string(runes))
	if err != nil {
		return 0
	}

	if isNegative {
		res *= -1
	}

	return checkint32(res)
}

func checkint32(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	return x
}
