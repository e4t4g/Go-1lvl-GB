package main

import (
	"fmt"
	"sort"
)

func main() {
	mas := make([]int, 10)

	for idx := range mas {
		fmt.Scan(&mas[idx])
	}

	fmt.Println(mas)

	sort.Ints(mas)

	fmt.Println(mas)
}
