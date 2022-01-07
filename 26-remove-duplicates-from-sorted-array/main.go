package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < lenNums-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == lenNums {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
