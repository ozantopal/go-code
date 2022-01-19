package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var profit int
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			tmp := prices[i] - min
			if tmp > profit {
				profit = tmp
			}
		}
	}

	return profit
}
