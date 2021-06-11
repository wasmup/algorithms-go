package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(OnesCountOn(0b0111_1111))    // 7
	fmt.Println(OnesCount(0b0111_1111))      // 7
	fmt.Println(bits.OnesCount(0b0111_1111)) // 7
}

func OnesCountOn(n uint) (count int) {
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return
}
func OnesCount(n uint) (count int) {
	for n != 0 {
		count++
		n &= n - 1
	}
	return
}
