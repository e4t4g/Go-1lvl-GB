package FibNumber

import (
	"fmt"
	"testing"
)

var result int

func TestFibonacci(t *testing.T) {
	var testSet = []struct{
		in int
		want int
	}{
		{0,0},
		{5,5},
		{8,21},
		{9,34},
	}
	for _,test := range testSet {
		if got := Fibonacci(test.in); got != test.want {
			t.Errorf("Got %d, want %d", got, test.want)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	var res int
	for i := 0; i<b.N; i++ {
		res = Fibonacci(15)
	}
	result = res
}

func ExampleFibonacci() {
	fmt.Println(Fibonacci(8))
	//Output: 21
}