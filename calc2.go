package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var sign string
	var a float64

	fmt.Print("Enter the first number: ")
	a, err := vali()
	if err != nil {
		fmt.Println("error NaN")
		return

	}

	fmt.Print("Enter one of the signs (+ - * /): ")
	fmt.Scan(&sign)

	fmt.Print("Enter the second number: ")
	b, err := vali()
	if err != nil {
		fmt.Println("error NaN")
		return
	}

	switch sign {
	case "+":
		fmt.Printf("Answer: %0.2f", a+b)
	case "-":
		fmt.Printf("Answer: %0.2f", a-b)
	case "*":
		fmt.Printf("Answer: %0.2f", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero.")
			break
		} else {
			fmt.Printf("Answer: %0.2f", a/b)
		}

	default:
		fmt.Println("Wrong")
	}

}

func vali() (float64, error) {
	var x string
	fmt.Scan(&x)
	y, err := strconv.ParseFloat(x,64)
	if err!=nil {
		 return y, errors.New("not empty error")
	} else {
		return y, nil
	}
}


