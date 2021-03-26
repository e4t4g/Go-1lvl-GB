//calculator

package main

import (
	"awesomeProject/calculator"
	"errors"
	"fmt"
	"strconv"
)

func EnterNumber() (float64, string, float64) {
	var A,B float64
	var Sign string

	fmt.Print("Enter the first number: ")
	A, err := Validation()
	if err != nil {
		fmt.Println("error NaN")

	}

	fmt.Print("Enter one of the signs (+ - * /): ")
	fmt.Scan(&Sign)

	fmt.Print("Enter the second number: ")
	B, err = Validation()
	if err != nil {
		fmt.Println("error NaN")

	}
	return A, Sign, B
}

func Validation() (float64, error) {
	var x string
	fmt.Scan(&x)
	y, err := strconv.ParseFloat(x,64)
	if err!=nil {
		return y, errors.New("not empty error")
	} else {
		return y, nil
	}
}


func main() {
	fmt.Println(calculator.Calc(EnterNumber()))
}

