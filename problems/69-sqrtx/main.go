package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low, high := 0, x
	var mid int
	for (high - low) > 1 {
		mid = low + (high-low)/2
		if mid == (x / mid) {
			return mid
		} else if mid > (x / mid) {
			high = mid
		} else if mid < (x / mid) {
			low = mid
		}
	}

	return low
}
