package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
	fmt.Println(hammingWeight(4294967293))
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}
