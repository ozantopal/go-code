package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}

func getRow(rowIndex int) []int {
	result := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}

		result = append(result, row)
	}

	return result[rowIndex]
}
