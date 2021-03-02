package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Println("Enter the slice`s capacity: ")
	fmt.Scan(&N)

	fmt.Println("Enter numbers of the slice: ")

	mas := make([]int, N)

	for idx := range mas {
		fmt.Scan(&mas[idx])
	}

	Sort(mas)

	fmt.Println(mas)

}

func Sort(mas []int) {
	for i := 1; i < len(mas); i++ {
		a := mas[i]
		k := i
		for ; k >= 1 && mas[k-1] > a; k-- {
			mas[k] = mas[k-1]
		}
		mas[k] = a
	}
}