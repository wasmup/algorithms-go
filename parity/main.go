package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(OnesCountOn(0b0111_1111))    // 7
	fmt.Println(OnesCount(0b0111_1111))      // 7
	fmt.Println(bits.OnesCount(0b0111_1111)) // 7

	fmt.Println(parity64(0b0111_1111)) // 1
}

func OnesCountOn(n uint) (count int) {
	for n != 0 {
		count += int(n & 1)
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

func parity64(n uint64) int {
	n ^= n >> 32
	n ^= n >> 16
	n ^= n >> 8
	n ^= n >> 4
	n ^= n >> 2
	n ^= n >> 1
	return int(n & 1)
}
