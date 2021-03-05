package main

import "fmt"

func main() {

	var n int
	_, err := fmt.Scan(&n)
	if err !=nil {
		fmt.Printf("error: %v", err)
	}

	var mapFib = map[int]int{}
	mapFib[n] = fibonacci(n)

	fmt.Print(mapFib)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}