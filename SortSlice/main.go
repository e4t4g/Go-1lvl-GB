package main

import (
	"Go-1lvl-GB/SortSlice/SortLogic"
	"fmt"
)

func main() {

	mas := make([]int, 10)

	for idx := range mas {
		fmt.Scan(&mas[idx])
	}

	fmt.Println(SortLogic.AutoSort(mas))
	fmt.Println(SortLogic.ManualSort(mas))
}

