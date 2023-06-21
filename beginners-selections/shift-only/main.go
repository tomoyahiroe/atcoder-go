package main

import (
	"fmt"
)

func main() {
	var N, count int
	fmt.Scanf("%d", &N)
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &array[i])
	}
	data := make(map[int]bool)
	data = mapEvenOdd(data, array)
	for isAllEven(data) {
		data = mapEvenOdd(data, divide(array))
		count = count + 1
	}
	fmt.Println(count)
}

func mapEvenOdd(data map[int]bool, array []int) map[int]bool {
	for i := range array {
		if array[i]%2 != 0 {
			data[array[i]] = false
		} else {
			data[array[i]] = true
		}
	}
	return data
}

func isAllEven(data map[int]bool) bool {
	for _, v := range data {
		if v == false {
			return false
		}
	}
	return true
}

func divide(array []int) []int {
	for i := range array {
		array[i] = array[i] / 2
	}
	return array
}
