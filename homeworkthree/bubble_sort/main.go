package main

import "fmt"

func main() {
	var myArr = []int64{1, 2, 3, 4, 5, 6, 7, 3, 2, 12, 4, 5, 34}
	var sortedArr = bubbleSort(myArr)
	for _, num := range sortedArr {
		fmt.Println(num)
	}
}

func bubbleSort(sliceToSort []int64) []int64 {
	result := make([]int64, len(sliceToSort))
	copy(result, sliceToSort)
	for i := 0; i < len(result)-1; i++ {
		for j := len(result) - 1; j > i; j-- {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
