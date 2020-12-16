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
	if num <= 1 {
		return num
	}

	if numberInMap, isContainNumber := fibonacciMap[num]; isContainNumber {
		return numberInMap
	} else {
		fibonacciMap[num] = FibonacciComputation(num-1) + FibonacciComputation(num-2)
	}
	return fibonacciMap[num]
}
