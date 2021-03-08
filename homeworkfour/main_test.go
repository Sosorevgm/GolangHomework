package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	name  string
	input int64
	want  int64
}{
	{
		name:  "Case №1",
		input: 1,
		want:  1,
	},
	{
		name:  "Case №2",
		input: 2,
		want:  1,
	},
	{
		name:  "Case №3",
		input: 10,
		want:  55,
	},
	{
		name:  "Case negative number",
		input: -10,
		want:  0,
	},
	{
		name:  "Case zero",
		input: 0,
		want:  0,
	},
}

func TestFibonacciComputation(t *testing.T) {
	for _, tc := range cases {
		getNumber := FibonacciComputation(tc.input)
		if getNumber != tc.want {
			t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, getNumber)
		}
	}
}

func TestFibonacciMapComputation(t *testing.T) {
	for _, tc := range cases {
		getNumber := FibonacciMapComputation(tc.input)
		if getNumber != tc.want {
			t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, getNumber)
		}
	}
}

func TestFibonacciLoopComputation(t *testing.T) {
	for _, tc := range cases {
		getNumber := FibonacciLoopComputation(tc.input)
		if getNumber != tc.want {
			t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, getNumber)
		}
	}
}

func ExampleFibonacciComputation() {
	fmt.Printf("%v", FibonacciLoopComputation(10))
	// Output: 55
}

func ExampleFibonacciMapComputation() {
	fmt.Printf("%v", FibonacciMapComputation(10))
	// Output: 55
}

func ExampleFibonacciLoopComputation() {
	fmt.Printf("%v", FibonacciLoopComputation(10))
	// Output: 55
}

func BenchmarkFibonacciComputation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciComputation(20)
	}
}

func BenchmarkFibonacciMapComputation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMapComputation(20)
	}
}

func BenchmarkFibonacciLoopComputation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciLoopComputation(20)
	}
}
