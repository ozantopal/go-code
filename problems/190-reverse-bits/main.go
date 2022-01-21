package main

import "fmt"

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = result<<1 | num&1
		num >>= 1
	}
	return result
}
