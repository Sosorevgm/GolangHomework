package main

import (
	"fmt"
)

var fibonacciMap = make(map[int64]int64)

func main() {
	fmt.Println(FibonacciComputation(6))
}

func FibonacciComputation(num int64) int64 {
	if num <= 1 {
		return num
	}

	if fibonacciMap[num] != 0 {
		return fibonacciMap[num]
	} else {
		fibonacciMap[num] = FibonacciComputation(num-1) + FibonacciComputation(num-2)
	}
	return fibonacciMap[num]
}
