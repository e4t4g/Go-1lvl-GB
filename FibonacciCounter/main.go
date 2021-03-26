package main

import (
	"awesomeProject/FibNumber"
	"fmt"
)

func main() {

	myMap := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	var a int


	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	if value, inMap := myMap[a]; inMap {
		fmt.Printf("MAP DATA: Fibonacci number for %v equal to: %v\n", a, value)

	} else {
		myMap[a] = FibNumber.Fibonacci(a)
		fmt.Printf("FRESH COUNT: Fibonacci number for %v equal to: %v\n", a, myMap[a])
	}


}
