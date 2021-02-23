package main

import "fmt"

func main() {
	var x int
	fmt.Print("Enter three digits number: ")
	fmt.Scanln(&x)
	if x < 1000 {

		a := x % 10
		b := ((x % 100) / 10)
		c := x / 100

		fmt.Printf("Hundreds: %v, Tens: %v, Units: %v", c, b, a)
	} else {
		fmt.Printf("%v contains more than three digits.", x)

	}

}

