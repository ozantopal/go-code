package main

import "fmt"

func main() {
	nums, target := []int{3, 2, 4}, 6

	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("Not found")
	}

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		val, isFound := m[target-num]
		if isFound {
			return []int{val, index}
		}
		
		m[num] = index
	}
	return nil
}