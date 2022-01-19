package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(4)
	obj.Push(3)
	obj.Push(1)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3)
	fmt.Println(param_4)
}

type MinStack struct {
	elements, min []int
	len           int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

func (ms *MinStack) Push(val int) {
	ms.elements = append(ms.elements, val)
	if ms.len == 0 {
		ms.min = append(ms.min, val)
	} else {
		min := ms.GetMin()
		if val < min {
			ms.min = append(ms.min, val)
		} else {
			ms.min = append(ms.min, min)
		}
	}
	ms.len++
}

func (ms *MinStack) Pop() {
	ms.len--
	ms.min = ms.min[:ms.len]
	ms.elements = ms.elements[:ms.len]
}

func (ms *MinStack) Top() int {
	return ms.elements[ms.len-1]
}

func (ms *MinStack) GetMin() int {
	return ms.min[ms.len-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
