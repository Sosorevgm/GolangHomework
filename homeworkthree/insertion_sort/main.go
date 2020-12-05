package main

import "fmt"

func main() {

	mArr := []int64{2}
	sortedSlice := insertionSort(mArr)
	for _, el := range sortedSlice {
		fmt.Println(el)
	}
}

func insertionSort(sliceToSort []int64) []int64 {
	sortedSlice := sliceToSort
	for i := 1; i < len(sliceToSort); i++ {
		currentElement := sortedSlice[i]
		previousKey := i - 1
		for previousKey >= 0 && sortedSlice[previousKey] > currentElement {
			sortedSlice[previousKey+1] = sortedSlice[previousKey]
			sortedSlice[previousKey] = currentElement
			previousKey--
		}
	}
	return sortedSlice
}
