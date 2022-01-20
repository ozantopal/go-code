package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(columnTitle string) int {
	result, val := 0, 0
	for i := 0; i < len(columnTitle); i++ {
		val = int(columnTitle[i] - 'A' + 1)
		result = result*26 + val
	}

	return result
}
