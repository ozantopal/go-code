package main

import "fmt"

func main() {
	n := 4
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second
		first, second = second, third
	}

	return second
}
