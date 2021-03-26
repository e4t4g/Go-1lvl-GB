package calculator

import (
	"fmt"
)

func Calc(A float64, Sign string, B float64) float64 {
	var z float64

	switch Sign {
	case "+":
		z = A + B
	case "-":
		z = A - B
	case "*":
		z = A * B
	case "/":
		if B == 0 {
			fmt.Println("Division by zero.")
			break
		} else {
			z = A / B
		}

	default:
		fmt.Println("Wrong")
	}
	return z
}
