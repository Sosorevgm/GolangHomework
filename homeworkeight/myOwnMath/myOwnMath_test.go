package myOwnMath

import (
	"fmt"
	"testing"
)

var numbers = []int32{
	1,
	2,
	3,
	4,
	200,
	142451,
}

var requiredNumbers = []int32{
	2,
	4,
	6,
	8,
	400,
	284902,
}

func TestMyOwnMathMethod(t *testing.T) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i]*2 != requiredNumbers[i] {
			t.Errorf("Multiplie %v * 2 = %v, want %v", numbers[i], numbers[i]*2, requiredNumbers[i])
		}
	}
}

func ExampleMyOwnMathMethod() {
	fmt.Printf("%v", MyOwnMathMethod(5))
	// Output: 10
}

func BenchmarkMyOwnMathMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyOwnMathMethod(5)
	}
}
