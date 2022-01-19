package main

import (
	"fmt"
	"math"
)

func main() {
	x := myAtoi("-123")
	fmt.Printf("%T | %v", x, x)
}

func myAtoi(str string) int {
	signs := []string{"+", "-"}
	var res int64 = 0
	var sign int64 = 1
	start := 0
	for id, ch := range str {
		if id == start && string(str[start]) == " " {
			start += 1
			continue
		}
		if id == start && string(ch) == signs[0] {
			sign = 1
			continue
		}
		if id == start && string(ch) == signs[1] {
			sign = -1
			continue
		}
		num, ok := getNumber(ch)
		if !ok {
			return int(res * sign)
		}
		if res*sign > math.MaxInt32/10 || (res*sign) == math.MaxInt32/10 && num > 7 {
			return math.MaxInt32
		}
		if res*sign < math.MinInt32/10 || (res*sign) == math.MinInt32/10 && num > 8 {
			return math.MinInt32
		}
		res = res*10 + int64(num)
	}
	return int(res * sign)

}

func getNumber(ch rune) (int, bool) {
	res := int(byte(ch) - "0"[0])
	if res > 10 || res < 0 {
		return res, false
	}
	return res, true
}
