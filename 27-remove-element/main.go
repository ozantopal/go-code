package main

import "fmt"

func main() {
	nums, val := []int{3, 2, 2, 3}, 3
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}

	result := 0
	for i := 0; i < lenNums; i++ {
		if nums[i] != val {
			if i != result {
				nums[i], nums[result] = nums[result], nums[i]
				result++
			} else {
				result++
			}
		}
	}

	return result
}
