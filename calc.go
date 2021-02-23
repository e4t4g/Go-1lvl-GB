package main

import "fmt"

func main() {

	var a,b float64
	var sign string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter one of the signs (+ - * /): ")
	fmt.Scan(&sign)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	switch sign {
	case "+":
		fmt.Printf("Answer: %0.2f", a+b)
	case "-":
		fmt.Printf("Answer: %0.2f", a-b)
	case "*":
		fmt.Printf("Answer: %0.2f", a*b)
	case "/":
		if b==0{
			fmt.Println("Division by zero.")
			break
		} else {
			fmt.Printf("Answer: %0.2f", a/b)
		}

	default:
		fmt.Println("Wrong")
	}

}

