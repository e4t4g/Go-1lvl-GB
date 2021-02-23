package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64

	fmt.Print("Enter area of the circle: ")
	fmt.Scanln(&S)

	fmt.Printf("Diameter: %0.2f;\n", (2 * math.Sqrt(S/math.Pi)));
	fmt.Printf("Circumference: %0.2f.", (2*math.Pi*(math.Sqrt(S/math.Pi))));
}

