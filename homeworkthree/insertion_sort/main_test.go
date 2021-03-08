package main

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var cases = []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "Case №1",
			input: []int64{2, 14, 3, 18, 2, 6},
			want:  []int64{2, 2, 3, 6, 14, 18},
		},
		{
			name:  "Case №2",
			input: []int64{-19, 1, 9, 0, 99, -102},
			want:  []int64{-102, -19, 0, 1, 9, 99},
		},
		{
			name:  "Case №3",
			input: []int64{1, 1, 1, 1, 2, 1, 1, 1},
			want:  []int64{1, 1, 1, 1, 1, 1, 1, 2},
		},
	}

	for _, tc := range cases {
		getSlice := insertionSort(tc.input)
		for i := 0; i < len(getSlice); i++ {
			if getSlice[i] != tc.want[i] {
				t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, getSlice)
				break
			}
		}
	}
}

func ExampleInsertionSort() {
	fmt.Printf("%v", insertionSort([]int64{3, 2, 1}))
	// Output: [1 2 3]
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort([]int64{12, 2, 14, 5, 4, 0, -12, 12, -14, 1144})
	}
}
