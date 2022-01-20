package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}

func convertToTitle(columnNumber int) string {
	result := []byte{}
	for columnNumber > 0 {
		result = append(result, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
