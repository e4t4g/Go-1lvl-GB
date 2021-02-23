package main

import "fmt"

func main() {
	var h,l int

	fmt.Print("Enter height of the rectangle: ")
	fmt.Scan(&h)

	fmt.Print("Enter length of the rectangle: ")
	fmt.Scan(&l)

	fmt.Println("S =", h*l)
}

