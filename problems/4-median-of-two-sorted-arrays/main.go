package main

import "fmt"

func main() {
	odd, even := []int{1, 2}, []int{3, 4}
	result := findMedianSortedArrays(odd, even)
	fmt.Println(result)
}

func findMedianSortedArrays(a []int, b []int) float64 {
	lenA, lenB := len(a), len(b)
	answer := make([]int, lenA+lenB)
	i, j, k := 0, 0, 0

	for k = 0; k < lenA+lenB; k++ {
		if i == lenA && j == lenB {
			break
		}
		if i == lenA {
			answer[k] = b[j]
			j++
			continue
		}
		if j == lenB {
			answer[k] = a[i]
			i++
			continue
		}

		if a[i] < b[j] {
			answer[k] = a[i]
			i++
		} else {
			answer[k] = b[j]
			j++
		}
	}

	var result float64
	if (lenA+lenB)%2 == 0 {
		result = float64((answer[(lenA+lenB)/2] + answer[((lenA+lenB)/2)-1])) / 2
	} else {
		result = float64(answer[int((lenA+lenB)/2)])
	}

	return result
}
