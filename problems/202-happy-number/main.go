package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	if n == 0 {
		return false
	}
	res := 0
	num := n
	record := map[int]int{}
	for {
		for num != 0 {
			res += (num % 10) * (num % 10)
			num = num / 10
		}
		if _, ok := record[res]; !ok {
			if res == 1 {
				return true
			}
			record[res] = res
			num = res
			res = 0
			continue
		} else {
			return false
		}
	}
}
