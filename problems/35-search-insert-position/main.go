package main

import "fmt"

func main() {
	nums, target := []int{1, 3, 5, 6}, 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	lowerThreshold, upperThreshold := 0, len(nums)-1
	for lowerThreshold <= upperThreshold {
		thresholdCandidate := lowerThreshold + (upperThreshold-lowerThreshold)>>1
		if nums[thresholdCandidate] >= target {
			upperThreshold = thresholdCandidate - 1
		} else {
			if (thresholdCandidate == len(nums)-1) || (nums[thresholdCandidate+1] >= target) {
				return thresholdCandidate + 1
			}
			lowerThreshold = thresholdCandidate + 1
		}
	}
	return 0
}
