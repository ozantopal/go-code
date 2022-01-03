package main

import (
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"

	fmt.Println(convert(s, 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	row, vertical := 0, true
	array := make([]string, numRows)
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		array[row] += string(s[i])
		if row == 0 {
			vertical = true
		}
		if row == numRows-1 {
			vertical = false
		}
		if vertical {
			row++
		} else {
			row--
		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		result += array[i]
	}
	return result
}
