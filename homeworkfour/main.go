package main

import (
	"fmt"
	"strconv"
)

var fibonacciMap = make(map[int64]int64)

func main() {
	fmt.Print("Введите целое число: ")

	var inputString string
	for {
		_, scanErr := fmt.Scanln(&inputString)
		if scanErr != nil {
			fmt.Println(scanErr)
			break
		}
		number, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Print("Введите целое число: ")
			continue
		}
		fmt.Println(FibonacciComputation(int64(number)))
		break
	}

}

func FibonacciComputation(num int64) int64 {
	if num < 0 {
		return 0
	}
	if num <= 1 {
		return num
	}
	return FibonacciComputation(num-1) + FibonacciComputation(num-2)
}

func FibonacciMapComputation(num int64) int64 {
	if num < 0 {
		return 0
	}
	if num <= 1 {
		return num
	}

	if numberInMap, isContainNumber := fibonacciMap[num]; isContainNumber {
		return numberInMap
	} else {
		fibonacciMap[num] = FibonacciMapComputation(num-1) + FibonacciMapComputation(num-2)
	}
	return fibonacciMap[num]
}

func FibonacciLoopComputation(num int64) int64 {
	var a int64 = 0
	var b int64 = 1
	for i := 0; i < int(num); i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}
