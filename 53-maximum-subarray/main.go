package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result, r, i := nums[0], 0, 0
	for i < len(nums) {
		r += nums[i]
		if r > result {
			result = r
		}
		if r < 0 {
			r = 0
		}
		i++
	}
	return result
}
