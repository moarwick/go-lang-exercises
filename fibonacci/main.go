package main

import (
	"fmt"
	"math/big"
)

func fib(nth int) *big.Int {
	if nth < 2 {
		return big.NewInt(int64(nth))
	}

	prev := big.NewInt(1)
	accum := big.NewInt(1)

	for i := 2; i <= nth; i++ {
		accum.Add(accum, prev)
		accum, prev = prev, accum
	}

	return accum
}

func main() {
	fmt.Println(0, fib(0))     // 0
	fmt.Println(1, fib(1))     // 1
	fmt.Println(2, fib(2))     // 1
	fmt.Println(3, fib(3))     // 2
	fmt.Println(4, fib(4))     // 3
	fmt.Println(5, fib(5))     // 5
	fmt.Println(6, fib(6))     // 8
	fmt.Println(10, fib(10))   // 55
	fmt.Println(80, fib(80))   // 23416728348467685
	fmt.Println(101, fib(101)) // 573147844013817084101
}
