package main

import "fmt"

func main() {

	myMap := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	var a int

	for i := 0; i < 15; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Printf("error: %v", err)
			break
		}
		if value, inMap := myMap[a]; inMap {
			fmt.Printf("MAP DATA: Fibonacci number for %v equal to: %v\n", a, value)

		} else {
			myMap[a] = fibonacci(a)
			fmt.Printf("FRESH COUNT: Fibonacci number for %v equal to: %v\n", a, myMap[a])
		}
	}

}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

