package SortLogic

import "sort"

func AutoSort(m []int) []int {
	sort.Ints(m)
	return m
}

func ManualSort(mas []int) []int {
	for i := 1; i < len(mas); i++ {
		a := mas[i]
		k := i
		for ; k >= 1 && mas[k-1] > a; k-- {
			mas[k] = mas[k-1]
		}
		mas[k] = a
	}
	return mas
}