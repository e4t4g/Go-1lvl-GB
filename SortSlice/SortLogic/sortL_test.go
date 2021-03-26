package SortLogic

import "testing"

var result []int
var ArrayTest = [] int {50,8,9,70,6}

func BenchmarkAutoSort(b *testing.B) {
	var res []int
	for i:=0; i<b.N; i++ {
		res = AutoSort(ArrayTest)
	}
	result = res
}

func BenchmarkManualSort(b *testing.B) {
	var res []int
	for i:=0; i<b.N; i++ {
		res = ManualSort(ArrayTest)
	}
	result = res
}